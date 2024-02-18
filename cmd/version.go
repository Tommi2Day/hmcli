// Package cmd commands
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version print version string",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			v := GetVersion(true)
			log.Debugf("Version: %s", v)
		},
	}
)

// Version, Build Commit and Date are filled in during build by the Makefile
// noinspection GoUnusedGlobalVariable
var (
	Name    = configName
	Version = "not set"
	Commit  = "snapshot"
	Date    = "undefined"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

// GetVersion extract compiled version info
func GetVersion(print bool) (txt string) {
	name := Name
	commit := Commit
	version := Version
	date := Date
	txt = fmt.Sprintf("%s version %s (%s - %s)", name, version, commit, date)
	if print {
		fmt.Println(txt)
	}
	return
}
