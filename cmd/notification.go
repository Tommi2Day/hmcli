package cmd

import (
	"fmt"
	"regexp"
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
	RunE:         getNotifications,
}

func init() {
	notificationsCmd.Flags().StringP("ignore", "I", "", "regexp to ignore notifications")
	RootCmd.AddCommand(notificationsCmd)
}

func getNotifications(cmd *cobra.Command, _ []string) error {
	var n hmlib.SystemNotificationResponse
	var err error
	var reIgnoreNotifications *regexp.Regexp

	log.Debug("notifications called")
	ignore, _ := cmd.Flags().GetString("ignore")
	if len(ignore) > 0 {
		log.Debugf("ignore parameter: %v", ignore)
		reIgnoreNotifications, err = regexp.Compile(ignore)
		if err != nil {
			return fmt.Errorf("cannot compile ignore regexp: %v", err)
		}
	}
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
	log.Debugf("ccu has %d notifications", l)
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
		if ok {
			nd.System = an.Name
			log.Debugf("system name set to %s", nd.System)
		}
		nd.Name, err = getNameFromStateList(id, e)
		if err != nil {
			l--
			log.Debugf("error getting name for notification: %v", err)
			p.Errors = append(p.Errors, err)
			continue
		}
		o := fmt.Sprintf("%s: %s(%s) since %s\n", nd.Type, nd.Name, nd.System, nd.Since.Format(time.RFC3339))
		if reIgnoreNotifications != nil && reIgnoreNotifications.MatchString(o) {
			log.Debugf("ignoring notification: %s", o)
			l--
			continue
		}
		longOutput += o
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

func getNameFromStateList(id string, e hmlib.Notification) (name string, err error) {
	// get state
	s, ok := hmlib.AllIDs[id]
	if !ok {
		err = fmt.Errorf("have no state for ID %s (%s)", id, e.Name)
		return
	}
	if s.Name != e.Name {
		err = fmt.Errorf("name mismatch for ID %s (%s)", id, e.Name)
		return
	}
	name = s.Name
	return
}
