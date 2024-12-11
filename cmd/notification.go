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
	notificationsCmd.Flags().BoolP("print", "p", false, "print ignored notifications")
	RootCmd.AddCommand(notificationsCmd)
}

const ignoredPrefix = "IGNORED: "

func getNotifications(cmd *cobra.Command, _ []string) error {
	log.Debug("notifications called")
	ignorePattern, _ := cmd.Flags().GetString("ignore")
	printIgnored, _ := cmd.Flags().GetBool("print")

	reIgnoreNotifications, err := setupIgnoreRegexp(ignorePattern)
	if err != nil {
		return err
	}
	if printIgnored {
		log.Debug("print ignored notifications enabled")
	}

	if _, err := hmlib.GetDeviceList("", false); err != nil {
		return err
	}

	notifications, err := hmlib.GetNotifications()
	if err != nil {
		return err
	}

	if len(notifications.Notifications) == 0 {
		NagiosResult(stateOK, "no notifications", "", nil)
		return nil
	}

	if _, err = hmlib.GetStateList(); err != nil {
		return err
	}

	return processNotifications(notifications, reIgnoreNotifications, printIgnored)
}

func setupIgnoreRegexp(ignorePattern string) (*regexp.Regexp, error) {
	if len(ignorePattern) == 0 {
		return nil, nil
	}
	log.Debugf("ignore parameter: %v", ignorePattern)
	reIgnore, err := regexp.Compile(ignorePattern)
	if err != nil {
		return nil, fmt.Errorf("cannot compile ignore regexp: %v", err)
	}
	return reIgnore, nil
}

func processNotifications(notifications hmlib.SystemNotificationResponse, reIgnoreNotifications *regexp.Regexp, printIgnored bool) error {
	notificationCount := len(notifications.Notifications)
	ignoredCount := 0
	notificationDetails := ""
	ignoredDetails := ""

	for _, notification := range notifications.Notifications {
		details, err := processNotificationDetail(notification)
		if err != nil {
			notificationCount--
			log.Debugf("error processing notification: %v", err)
			plugin.Errors = append(plugin.Errors, err)
			continue
		}

		output := fmt.Sprintf("%s: %s(%s) since %s\n", details.Type, details.Name, details.System, details.Since.Format(time.RFC3339))
		if reIgnoreNotifications != nil && reIgnoreNotifications.MatchString(output) {
			log.Debugf("ignoring notification: %s", output)
			notificationCount--
			if !printIgnored {
				continue
			}
			ignoredCount++
			ignoredDetails += output
			output = ""
		}
		notificationDetails += output
	}

	finalOutput := fmt.Sprintf("%d notifications pending", notificationCount)
	if ignoredCount > 0 {
		finalOutput += fmt.Sprintf(", %d ignored", ignoredCount)
		notificationDetails += fmt.Sprintf("%s\n%s", ignoredPrefix, ignoredDetails)
	}

	perfdata := []nagios.PerformanceData{
		{
			Label: "notifications",
			Value: fmt.Sprintf("%d", notificationCount),
			Warn:  hmWarnThreshold,
			Crit:  hmCritThreshold,
		},
	}

	NagiosResult(stateOK, finalOutput, notificationDetails, perfdata)
	return nil
}

func processNotificationDetail(notification hmlib.Notification) (*hmlib.NotificationDetail, error) {
	ts, _ := common.GetInt64Val(notification.Timestamp)
	notificationTime := time.Unix(ts, 0)

	addressParts := strings.Split(notification.Name, ".")
	if len(addressParts) < 2 {
		return nil, fmt.Errorf("invalid notification name format: %s", notification.Name)
	}

	address := strings.Split(addressParts[1], ":")[0]
	system := "unknown"
	if device, exists := hmlib.DeviceAddressMap[address]; exists {
		system = device.Name
		log.Debugf("system name set to %s", system)
	}

	name, err := getNameFromStateList(notification.IseID, notification)
	if err != nil {
		return nil, err
	}

	return &hmlib.NotificationDetail{
		Since:   notificationTime,
		Address: address,
		System:  system,
		Type:    notification.Type,
		Name:    name,
	}, nil
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
