package cmd

import (
	"fmt"
	"regexp"

	"github.com/tommi2day/gomodules/common"

	"github.com/atc0005/go-nagios"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tommi2day/gomodules/hmlib"
)

var datapointCmd = &cobra.Command{
	Use:   "datapoint",
	Short: "command related to datapoints",
}
var datapointListCmd = &cobra.Command{
	Use:          "list",
	Short:        "List all data points for a device",
	Long:         ``,
	SilenceUsage: true,
	RunE:         datapointList,
}

var datapointCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check a single datapoint",
	Long: `select a single datapoint by name or id and check its value
You can set -w or -c to set the warning or critical threshold on numeric values`,
	SilenceUsage: false,
	RunE:         datapointCheck,
}

var stateListResponse hmlib.StateListResponse

func init() {
	datapointListCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		_ = command.Flags().MarkHidden("warn")
		_ = command.Flags().MarkHidden("crit")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
	datapointListCmd.Flags().StringP("match", "m", "", "list data points with name matching regexp")
	datapointCheckCmd.Flags().StringP("name", "n", "", "datapoint name(one) to query")
	datapointCheckCmd.Flags().StringP("id", "i", "", "datapoint id (one) to query")
	datapointCheckCmd.Flags().StringP("match", "m", "", "datapoint value must match this regexp")
	RootCmd.AddCommand(datapointCmd)
	datapointCmd.AddCommand(datapointListCmd)
	datapointCmd.AddCommand(datapointCheckCmd)
}

func datapointList(cmd *cobra.Command, _ []string) error {
	var err error
	m, _ := cmd.Flags().GetString("match")
	log.Debugf("match parameter: %v", m)
	stateListResponse, err = hmlib.GetStateList()
	sl := len(stateListResponse.StateDevices)
	if err != nil {
		err = fmt.Errorf("cannot get state list: %v", err)
		return err
	}
	if sl == 0 {
		err = fmt.Errorf("no devices found in state list")
		return err
	}
	log.Debugf("found %d devices in state list", sl)

	// print data points
	count, err := printDatapointList(m)
	if err != nil {
		return err
	}
	if count == 0 && m != "" {
		err = fmt.Errorf("no matching data points found")
		return err
	}
	log.Debugf("SUCCESS: %d data points", count)
	return nil
}
func printDatapointList(m string) (count int, err error) {
	var re *regexp.Regexp
	if m != "" {
		re, err = regexp.Compile(m)
		if err != nil {
			err = fmt.Errorf("cannot compile regexp %s: %v", m, err)
			return
		}
	}
	for _, d := range stateListResponse.StateDevices {
		for _, c := range d.Channels {
			for _, dp := range c.Datapoints {
				if m != "" && !re.MatchString(dp.Name) {
					continue
				}
				count++
				out := fmt.Sprintf("Device: %s Channel: %s Datapoint: %s(%s) = %s %s",
					d.Name, c.Name, dp.Name, dp.IseID, dp.Value, dp.ValueUnit)
				fmt.Println(out)
			}
		}
	}
	return
}
func validateDatapointCheckInput(i, n, m string) (dp hmlib.Datapoint, err error) {
	// check regexp
	if m != "" {
		_, err = regexp.Compile(m)
		if err != nil {
			err = fmt.Errorf("cannot compile regexp %s: %v", m, err)
			return
		}
	}
	if n == "" && i == "" {
		err = fmt.Errorf("please provide a name or id")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if n != "" && i != "" {
		err = fmt.Errorf("please provide only a name or id")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if n != "" {
		x, ok := hmlib.NameIDMap[n]
		if !ok {
			err = fmt.Errorf("datapoint name %s not found", n)
			log.Debugf("UNKNOWN: %s", err.Error())
			return
		}
		i = x
	}
	if i == "" {
		err = fmt.Errorf("please provide a valid name or id")
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	e, ok := hmlib.AllIDs[i]
	if !ok {
		err = fmt.Errorf("datapoint id %s not found", i)
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}
	if e.EntryType != "Datapoint" {
		err = fmt.Errorf("id %s is not a datapoint", i)
		log.Debugf("UNKNOWN: %s", err.Error())
		return
	}

	// cast datapoint back
	dp = e.Entry.(hmlib.Datapoint)
	return
}
func datapointCheck(cmd *cobra.Command, _ []string) error {
	var err error
	var re *regexp.Regexp
	var dp hmlib.Datapoint
	n, _ := cmd.Flags().GetString("name")
	i, _ := cmd.Flags().GetString("id")
	m, _ := cmd.Flags().GetString("match")
	log.Debugf("name parameter: %v", n)
	log.Debugf("id parameter: %v", i)
	log.Debugf("match parameter: %v", m)
	stateListResponse, err = hmlib.GetStateList()
	sl := len(stateListResponse.StateDevices)
	if err != nil {
		err = fmt.Errorf("cannot get state list: %v", err)
		return err
	}
	if sl == 0 {
		err = fmt.Errorf("no devices found in state list")
		return err
	}
	log.Debugf("found %d devices in state list", sl)

	// validate and get datapoint
	dp, err = validateDatapointCheckInput(i, n, m)
	if err != nil {
		return err
	}

	longOutput := fmt.Sprintf("Datapoint:%s ID:%s Value:=%s %s", dp.Name, dp.IseID, dp.Value, dp.ValueUnit)
	output := fmt.Sprintf("%s=%s", dp.Name, dp.Value)

	// prepare performance data only for numeric values
	preparePD := false
	var performanceData []nagios.PerformanceData
	switch dp.ValueType {
	case "2":
		// boolean
	case "4":
		// float
		preparePD = common.IsNumeric(dp.Value)
	case "16":
		// integer
		preparePD = common.IsNumeric(dp.Value)
	case "20":
		// text
	}
	log.Debugf("value is numeric: %v", preparePD)
	if preparePD {
		performanceData = []nagios.PerformanceData{{Label: fmt.Sprintf("%s(%s)", dp.Name, dp.IseID),
			Value: dp.Value,
			Warn:  hmWarnThreshold,
			Crit:  hmCritThreshold}}
		log.Debugf("performance data: %v", performanceData)
	}
	// check match value
	if m != "" {
		re = regexp.MustCompile(m)
		if !re.MatchString(dp.Value) {
			output = fmt.Sprintf("CRITICAL: %s returned value '%s', does not match '%s'", dp.Name, dp.Value, m)
			log.Debug(output)
			NagiosResult("CRITICAL", output, longOutput, performanceData)
			return nil
		}
		log.Debugf("value %s matches regexp %s", dp.Value, m)
	}
	// set final nagios state
	NagiosResult("OK", output, longOutput, performanceData)
	return nil
}
