package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/atc0005/go-nagios"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

var notificationsCmd = &cobra.Command{
	Use:     "notifications",
	Aliases: []string{"notification"},
	Short:   "check hmWarnThreshold notifications",
	Long: `List all current notifications. 
You can set -w or -c to set the warning or critical threshold on notification count`,
	SilenceUsage: true,
	RunE:         notifications,
}

func init() {
	RootCmd.AddCommand(notificationsCmd)
}

func notifications(_ *cobra.Command, _ []string) error {
	var n hmlib.SystemNotificationResponse
	var err error

	log.Debug("notifications called")
	p := GetHmPlugin()
	_, err = hmlib.GetDeviceList("", false)
	if err != nil {
		return err
	}
	n, err = hmlib.GetNotifications()
	if err != nil {
		return err
	}
	l := len(n.Notifications)
	if l == 0 {
		NagiosResult("OK", "no notifications", "", nil)
		return nil
	}
	_, err = hmlib.GetStateList()
	if err != nil {
		return err
	}
	longOutput := ""
	// iterate over all notifications
	for _, e := range n.Notifications {
		var nd hmlib.NotificationDetail
		id := e.IseID
		ts, _ := common.GetInt64Val(e.Timestamp)
		nd.Since = time.Unix(ts, 0)
		nd.Type = e.Type
		a := strings.Split(e.Name, ".")
		nd.System = "unknown"
		aa := strings.Split(a[1], ":")
		nd.Address = aa[0]
		log.Debugf("address: %s", nd.Address)
		an, ok := hmlib.DeviceAddressMap[nd.Address]
		log.Debugf("an: %v", an)
		if ok {
			nd.System = an.Name
		}
		// get state
		s, ok := hmlib.AllIDs[id]
		if !ok {
			pe := fmt.Errorf("have no state for ID %s (%s)", id, e.Name)
			l--
			p.Errors = append(p.Errors, pe)
			continue
		}
		if s.Name != e.Name {
			pe := fmt.Errorf("name mismatch for ID %s (%s)", id, e.Name)
			l--
			p.Errors = append(p.Errors, pe)
			continue
		}
		nd.Name = s.Name
		longOutput += fmt.Sprintf("%s: %s(%s) since %s\n", nd.Type, nd.Name, nd.System, nd.Since.Format(time.RFC3339))
	}

	// set final nagios state
	output := fmt.Sprintf("%d notifications pending", l)
	perfdata := []nagios.PerformanceData{{Label: "notifications",
		Value: fmt.Sprintf("%d", l),
		Warn:  hmWarnThreshold,
		Crit:  hmCritThreshold}}

	NagiosResult("OK", output, longOutput, perfdata)
	return nil
}
