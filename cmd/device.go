package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tommi2day/gomodules/hmlib"
)

var devicesCmd = &cobra.Command{
	Use:     "device",
	Aliases: []string{"devices"},
	Short:   "command related to devices",
}
var deviceListCmd = &cobra.Command{
	Use:          "list",
	Short:        "list devices",
	Long:         `list all devices or a specify one`,
	SilenceUsage: true,
	RunE:         deviceList,
}

func init() {
	deviceListCmd.Flags().StringP("address", "a", "", "device addresses to query, comma separated")
	deviceListCmd.Flags().StringP("name", "n", "", "device name to query, comma separated")
	deviceListCmd.Flags().StringP("id", "i", "", "device ids to query, comma separated")
	deviceListCmd.Flags().Bool("internal", false, "also list internal channels")
	devicesCmd.AddCommand(deviceListCmd)
	RootCmd.AddCommand(devicesCmd)
}

func getDevicesByName(name string) (devices []hmlib.DeviceListEntry) {
	for _, e := range hmlib.DeviceIDMap {
		if e.Name == name {
			devices = append(devices, e)
		}
	}
	return
}
func deviceList(cmd *cobra.Command, _ []string) error {
	log.Debug("deviceList called")
	n, _ := cmd.Flags().GetString("name")
	a, _ := cmd.Flags().GetString("address")
	i, _ := cmd.Flags().GetString("id")
	all := cmd.Flags().Changed("internal")
	log.Debugf("address parameter: %v", a)
	log.Debugf("name parameter: %v", n)
	log.Debugf("id parameter: %v", i)
	_, err := hmlib.GetDeviceList("", all)
	if err != nil {
		return err
	}
	// list all devices, no filter given
	if len(i) == 0 && len(a) == 0 && len(n) == 0 {
		for _, e := range hmlib.DeviceIDMap {
			fmt.Println(e)
		}
		return nil
	}
	// query by name
	if len(n) > 0 {
		entries := getDevicesByName(n)
		for _, entry := range entries {
			log.Debugf("found device %s by name", n)
			fmt.Println(entry)
		}
		if len(entries) == 0 {
			fmt.Printf("device with name %s not found\n", n)
		}
	}
	// query by address
	if len(a) > 0 {
		printEntry(a, hmlib.DeviceAddressMap)
		return nil
	}
	// query by id
	if len(i) > 0 {
		printEntry(i, hmlib.DeviceIDMap)
	}
	// list all deviceList
	return nil
}

func printEntry(id string, hmmap map[string]hmlib.DeviceListEntry) {
	entry, ok := hmmap[id]
	if !ok {
		fmt.Printf("device with ID %s not found\n", id)
	} else {
		log.Debugf("found device %s", entry.Name)
		fmt.Println(entry)
	}
}
