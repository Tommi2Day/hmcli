// Package cmd commands
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/tommi2day/gomodules/hmlib"

	"github.com/tommi2day/gomodules/common"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	cfgFile         string
	debugFlag       = false
	unitTestFlag    = false
	hmToken         string
	hmURL           string
	hmWarnThreshold string
	hmCritThreshold string

	// RootCmd entry point to start
	RootCmd = &cobra.Command{
		Use:           "check_hm",
		Short:         "check_hm – Check Homematic Health Monitor",
		Long:          `Nagios/Icinga plugin  check Homematic/Raspberrymatic status with XMLAPI`,
		SilenceErrors: true,
	}
)

const (
	// allows you to override any config values using
	// env APP_MY_VAR = "MY_VALUE"
	// e.g. export APP_LDAP_USERNAME test
	// maps to ldap.username
	configEnvPrefix = "CHECK_HM"
	configName      = "check_hm"
	configType      = "yaml"
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "verbose debug output")
	RootCmd.PersistentFlags().BoolVarP(&unitTestFlag, "unit-test", "", false, "redirect output for unit tests")
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "C", "", "config file name")
	RootCmd.PersistentFlags().StringVarP(&hmToken, "token", "t", "", "Homematic XMLAPI Token")
	RootCmd.PersistentFlags().StringVarP(&hmURL, "url", "u", "https://ccu", "Homematic URL (default: https://ccu)")
	RootCmd.PersistentFlags().StringVarP(&hmWarnThreshold, "warn", "w", "", "warning level")
	RootCmd.PersistentFlags().StringVarP(&hmCritThreshold, "crit", "c", "", "critical level")
	// don't have variables populated here
	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}

// Execute run application
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		p := GetHmPlugin()
		p.Errors = append(p.Errors, err)
		log.Debugf("return UNKNOWN, errors: %v", p.Errors)
		NagiosResult("UNKNOWN", "", "", nil)
	}
}

func initConfig() {
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.SetConfigFile(configName + "." + configType)
	if cfgFile == "" {
		// Search config in $HOME/etc and current directory.
		viper.AddConfigPath(".")
	} else {
		// set filename form cli
		viper.SetConfigFile(cfgFile)
	}

	// env var overrides
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(configEnvPrefix)
	// env var `LDAP_USERNAME` will be mapped to `ldap.username`
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	haveConfig, err := processConfig()

	// check flags
	processFlags()

	// logger settings
	log.SetLevel(log.ErrorLevel)
	if debugFlag {
		// report function name
		log.SetReportCaller(true)
		log.SetLevel(log.DebugLevel)
	}

	logFormatter := &prefixed.TextFormatter{
		DisableColors:   unitTestFlag,
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123,
	}
	log.SetFormatter(logFormatter)

	if unitTestFlag {
		log.SetOutput(RootCmd.OutOrStdout())
	}
	// debug config file
	if haveConfig {
		log.Debugf("found configfile %s", cfgFile)
	} else {
		log.Debugf("Error using %s config: %s", configType, err)
	}

	// validate method
}

// processConfig reads in config file and ENV variables if set.
func processConfig() (bool, error) {
	err := viper.ReadInConfig()
	haveConfig := false
	if err == nil {
		cfgFile = viper.ConfigFileUsed()
		haveConfig = true
	}
	return haveConfig, err
}

func processFlags() {
	if common.CmdFlagChanged(RootCmd, "debug") {
		viper.Set("debug", debugFlag)
	}
	if common.CmdFlagChanged(RootCmd, "token") {
		viper.Set("token", hmToken)
	}
	if common.CmdFlagChanged(RootCmd, "url") {
		viper.Set("url", hmURL)
	}
	debugFlag = viper.GetBool("debug")
	hmToken = viper.GetString("token")
	hmURL = viper.GetString("url")
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)
}
