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

var sysvarsCmd = &cobra.Command{
	Use:   "sysvar",
	Short: "command related to system variables",
}
var sysvarListCmd = &cobra.Command{
	Use:          "list",
	Short:        "List all system variables",
	Long:         ``,
	SilenceUsage: true,
	RunE:         sysvarList,
}

var sysvarCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check a single system variable",
	Long: `select a single system variable by name or id and check its value
You can set -w or -c to set the warning or critical threshold on numeric values`,
	SilenceUsage: false,
	RunE:         sysvarCheck,
}

func init() {
	sysvarListCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		_ = command.PersistentFlags().MarkHidden("warn")
		_ = command.PersistentFlags().MarkHidden("crit")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
	sysvarCheckCmd.Flags().StringP("name", "n", "", "value name(one) to query")
	sysvarCheckCmd.Flags().StringP("id", "i", "", "variable id (one) to query")
	sysvarCheckCmd.Flags().StringP("match", "m", "", "value must match this regexp")
	sysvarsCmd.AddCommand(sysvarListCmd)
	sysvarsCmd.AddCommand(sysvarCheckCmd)
	RootCmd.AddCommand(sysvarsCmd)
}

func sysvarList(cmd *cobra.Command, _ []string) error {
	log.Debug("sysvarList called")
	n, _ := cmd.Flags().GetString("name")
	i, _ := cmd.Flags().GetString("id")
	log.Debugf("name parameter: %v", n)
	log.Debugf("id parameter: %v", i)
	err := hmlib.GetSysvarList(true)
	if err != nil {
		return err
	}
	// print
	if len(hmlib.SysVarIDMap) > 0 {
		printSysVarList(n, i)
		return nil
	}
	err = fmt.Errorf("no sysvars found")
	return err
}
func printSysVarList(n string, i string) {
	for _, e := range hmlib.SysVarIDMap {
		if n == "" && i == "" {
			fmt.Println(e)
		}
		if n != "" && e.Name == n {
			fmt.Println(e)
			return
		}
		if i != "" && e.IseID == i {
			fmt.Println(e)
			return
		}
	}
}
func sysvarCheck(cmd *cobra.Command, _ []string) error {
	log.Debug("sysvar check called")
	var err error
	n, _ := cmd.Flags().GetString("name")
	i, _ := cmd.Flags().GetString("id")
	m, _ := cmd.Flags().GetString("match")
	log.Debugf("name parameter: %v", n)
	log.Debugf("id parameter: %v", i)
	if n == "" && i == "" {
		return fmt.Errorf("please provide a name or id")
	}
	if n != "" && i != "" {
		return fmt.Errorf("please provide only a name or id")
	}

	// get list
	entry, err := getSysVar(n, i)
	if err != nil {
		return fmt.Errorf("UNKNOWN: %s", err)
	}

	// get formatted value of sysvar
	v := entry.GetValue()
	longOutput := entry.String()
	output := fmt.Sprintf("%s=%s (%s)", entry.Name, v, entry.Unit)

	// prepare performance data if value is numeric
	preparePD := common.IsNumeric(v)
	log.Debugf("value is numeric: %v", preparePD)
	var performanceData []nagios.PerformanceData
	if preparePD {
		performanceData = []nagios.PerformanceData{{Label: fmt.Sprintf("%s(%s).%s", entry.Name, entry.IseID, n),
			Value: v,
			Warn:  hmWarnThreshold,
			Crit:  hmCritThreshold}}
		log.Debugf("performance data: %v", performanceData)
	}
	status := stateOK
	if m != "" {
		status, output = checkMatch(entry.Name, v, m)
	}
	// set final nagios state
	NagiosResult(status, output, longOutput, performanceData)
	return nil
}
func getSysVar(n string, i string) (entry hmlib.SysVarEntry, err error) {
	err = hmlib.GetSysvarList(false)
	if err != nil {
		err = fmt.Errorf("getSysvarList failed:%s", err)
		log.Debugf(err.Error())
		return
	}
	// query by name
	found := false
	if n != "" {
		for _, e := range hmlib.SysVarIDMap {
			if e.Name == n {
				i = e.IseID
				log.Debugf("found id %s for name %s", i, n)
				found = true
				break
			}
		}
		if !found {
			err = fmt.Errorf("sysvar with name %s not found", n)
			log.Debugf(err.Error())
			return
		}
	}
	entry, ok := hmlib.SysVarIDMap[i]
	if !ok {
		err = fmt.Errorf("UNKNOWN:sysvar id %s not found", i)
		log.Debugf(err.Error())
		return
	}
	return
}
func checkMatch(name string, v string, m string) (status string, output string) {
	// check match value
	status = "OK"
	re, err := regexp.Compile(m)
	if err != nil {
		output = fmt.Sprintf("UNKNOWN: regexp compile %s failed: %s", m, err)
		log.Debugf(output)
		status = stateUnknown
		return
	}
	if !re.MatchString(v) {
		output = fmt.Sprintf("CRITICAL: sysvar %s value %s, does not match %s", name, v, m)
		log.Debugf(output)
		status = stateCritical
		return
	}
	log.Debugf("value %s matches regexp %s", v, m)
	return
}
