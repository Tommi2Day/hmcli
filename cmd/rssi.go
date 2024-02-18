package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tommi2day/gomodules/hmlib"
)

var rssiCmd = &cobra.Command{
	Use:          "rssi",
	Short:        "list rssi values",
	Long:         `List available rssi values of devices`,
	SilenceUsage: true,
	RunE:         rssilist,
}

func init() {
	rssiCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		_ = command.Flags().MarkHidden("warn")
		_ = command.Flags().MarkHidden("critical")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
	RootCmd.AddCommand(rssiCmd)
}

func rssilist(_ *cobra.Command, _ []string) (err error) {
	log.Debug("rssilist called")
	var result hmlib.RssiListResponse
	parameter := map[string]string{}
	err = hmlib.QueryAPI(hmlib.RssiEndpoint, &result, parameter)
	if err == nil {
		fmt.Println(result)
	}
	return
}
