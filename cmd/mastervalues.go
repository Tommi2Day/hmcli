package cmd

import (
	"fmt"
	"regexp"

	"github.com/atc0005/go-nagios"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

var valueCmd = &cobra.Command{
	Use:     "value",
	Aliases: []string{"values", "mastervalue", "mastervalues"},
	Short:   "command related to device master values",
}
var valueListCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long: `List for a given device all mastervalues  or a named value. 
Address or id must be given  `,
	SilenceUsage: false,
	RunE:         mastervalueList,
}

var valueCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check a single device Value",
	Long: `this retrives a given device mastervalue and may check if it contains a string (-m)
You can set -w or -c to set the warning or critical threshold on numeric values`,
	SilenceUsage: false,
	RunE:         mastervalueCheck,
}

func init() {
	valueListCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		_ = command.Flags().MarkHidden("warn")
		_ = command.Flags().MarkHidden("crit")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
	valueListCmd.Flags().StringP("address", "a", "", "device addresses to query, comma separated")
	valueListCmd.Flags().StringP("id", "i", "", "device ids to query, comma separated")
	valueListCmd.Flags().StringP("name", "n", "", "requested mastervalue names, comma separated")
	valueCheckCmd.Flags().StringP("address", "a", "", "device address (only one) to query, alternative to id")
	valueCheckCmd.Flags().StringP("id", "i", "", "device id (one) to query, alternative to address")
	valueCheckCmd.Flags().StringP("name", "n", "", "requested master value (not device) name ( only one)")
	valueCheckCmd.Flags().StringP("match", "m", "", "returned master value must match this regexp")
	RootCmd.AddCommand(valueCmd)
	valueCmd.AddCommand(valueListCmd)
	valueCmd.AddCommand(valueCheckCmd)
}

func mastervalueList(cmd *cobra.Command, _ []string) error {
	log.Debug("value list called")
	a, _ := cmd.Flags().GetString("address")
	i, _ := cmd.Flags().GetString("id")
	n, _ := cmd.Flags().GetString("name")
	log.Debugf("address parameter: %v", a)
	log.Debugf("id parameter: %v", i)
	log.Debugf("name parameter: %v", n)
	if a == "" && i == "" && n == "" {
		err := fmt.Errorf("please provide an address or id or exact name")
		log.Debugf(err.Error())
		return err
	}
	// get device list
	// todo: limit device ids, if given (use i instead of "")
	// But currently it returns invalid xml(https://github.com/homematic-community/XML-API/issues/93)
	_, err := hmlib.GetDeviceList("", false)
	if err != nil {
		return err
	}

	// get id from address if this given
	if len(a) > 0 {
		entry, ok := hmlib.DeviceAddressMap[a]
		if !ok {
			fmt.Printf("device with address %s not found\n", a)
		} else {
			i = entry.IseID
		}
	}

	// get values
	mv, err := hmlib.GetMasterValues(i, n)
	if err != nil {
		return err
	}
	// print
	l := len(mv.MasterValueDevices)
	log.Debugf("found %d devices", l)
	if l > 0 {
		for _, e := range mv.MasterValueDevices {
			fmt.Printf("Device (ID=%s): %s, Type: %s\n", e.IseID, e.Name, e.DeviceType)
			for _, v := range e.MasterValue {
				fmt.Printf("  %s=%s\n", v.Name, v.Value)
			}
		}
		return nil
	}
	err = fmt.Errorf("no devices found")
	return err
}

func mastervalueCheckInput(a, i, n, m string) (id string, err error) {
	log.Debugf("name parameter: %v", n)
	log.Debugf("id parameter: %v", i)
	log.Debugf("match parameter: %v", m)
	log.Debugf("address parameter: %v", a)
	id = i
	if a == "" && i == "" {
		err = fmt.Errorf("please provide an address or id")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if a != "" && i != "" {
		err = fmt.Errorf("UNKNOWN: please provide only an address or id")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if a != "" {
		_, err = hmlib.GetDeviceList("", false)
		if err != nil {
			return
		}

		entry, ok := hmlib.DeviceAddressMap[a]
		if !ok {
			err = fmt.Errorf("device address %s not found", a)
			log.Debugf("UNKNOWN: %s", err.Error())
			return
		}
		i = entry.IseID
	}
	if n == "" && i == "" {
		err = fmt.Errorf("value name or id not provided")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	log.Debugf("use id parameter: %s", i)
	// check regexp
	if m != "" {
		_, err = regexp.Compile(m)
		if err != nil {
			err = fmt.Errorf("cannot compile regexp %s: %v", m, err)
			return
		}
	}
	return
}

func getNamedMasterValue(n string, d hmlib.MasterValueDevice) (v string, err error) {
	// find named value
	found := false
	if n != "" {
		for _, e := range d.MasterValue {
			if e.Name == n {
				v = e.Value
				log.Debugf("found value %s for name %s", v, n)
				found = true
				break
			}
		}
		if !found {
			err = fmt.Errorf("value with name %s not found for device %s", n, d.Name)
			log.Debugf("UNKNOWN: %s", err.Error())
			return
		}
	}
	return
}

func getMasterValueDevice(i string, n string) (d hmlib.MasterValueDevice, err error) {
	// get list
	mv, err := hmlib.GetMasterValues(i, n)
	if err != nil {
		err = fmt.Errorf("getMasterValueList failed:%s", err)
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if len(mv.MasterValueDevices) == 0 {
		err = fmt.Errorf("device not found")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	// find device
	found := false
	for _, e := range mv.MasterValueDevices {
		if e.IseID == i {
			log.Debugf("found device %s in response", i)
			d = e
			found = true
			break
		}
	}
	if !found {
		err = fmt.Errorf("device with id %s not in response", i)
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	return
}

func mastervalueCheck(cmd *cobra.Command, _ []string) error {
	log.Debug("value check called")
	var err error
	var re *regexp.Regexp
	var d hmlib.MasterValueDevice
	v := ""

	a, _ := cmd.Flags().GetString("address")
	n, _ := cmd.Flags().GetString("name")
	i, _ := cmd.Flags().GetString("id")
	m, _ := cmd.Flags().GetString("match")
	i, err = mastervalueCheckInput(a, i, n, m)
	if err != nil {
		return err
	}

	// get device by id
	d, err = getMasterValueDevice(i, n)
	if err != nil {
		return err
	}

	// get value by name
	v, err = getNamedMasterValue(n, d)
	if err != nil {
		log.Debugf("UNKNOWN: %s", err.Error())
		return err
	}

	// prepare output
	longOutput := fmt.Sprintf("Device:%s ID:%s ValueName:%s=%s", d.IseID, d.Name, n, v)
	output := fmt.Sprintf("%s=%s", n, v)

	// prepare performance data only for numeric values
	preparePD := common.IsNumeric(v)
	log.Debugf("value is numeric: %v", preparePD)
	var performanceData []nagios.PerformanceData
	if preparePD {
		performanceData = []nagios.PerformanceData{{Label: fmt.Sprintf("%s(%s).%s", d.Name, d.IseID, n),
			Value: v,
			Warn:  hmWarnThreshold,
			Crit:  hmCritThreshold}}
		log.Debugf("performance data: %v", performanceData)
	}

	// check match value
	if m != "" {
		re = regexp.MustCompile(m)
		if !re.MatchString(v) {
			output = fmt.Sprintf("CRITICAL: %s.%s returned value '%s', does not match '%s'", d.Name, n, v, m)
			log.Debugf(output)
			NagiosResult("CRITICAL", output, longOutput, performanceData)
			return nil
		}
		log.Debugf("value %s matches regexp %s", v, m)
	}
	// set final nagios state
	NagiosResult("OK", output, longOutput, performanceData)
	return nil
}
