package cmd

import (
	"testing"

	"github.com/tommi2day/check_hm/test"

	"github.com/atc0005/go-nagios"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

func TestNotifications(t *testing.T) {
	var httpClient = resty.New()
	test.Testinit(t)
	httpmock.ActivateNonDefault(httpClient.GetClient())
	defer httpmock.DeactivateAndReset()
	hmlib.SetHTTPClient(httpClient)
	defer httpmock.DeactivateAndReset()
	hmToken = MockToken
	hmURL = MockURL
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)

	// mock the response for notifications
	responderURL := hmURL + hmlib.NotificationsEndpoint
	httpmock.RegisterResponder("GET", responderURL, httpmock.NewStringResponder(200, NotificationsTest))

	stateListURL := hmURL + hmlib.StateListEndpoint
	httpmock.RegisterResponder("GET", stateListURL, httpmock.NewStringResponder(200, StateListTest))

	deviceListURL := hmURL + hmlib.DeviceListEndpoint
	httpmock.RegisterResponder("GET", deviceListURL, httpmock.NewStringResponder(200, DeviceListTest))

	t.Run("notifications cmd", func(t *testing.T) {
		args := []string{
			"notifications",
			"--debug",
			"--warn", "0",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		SetHmPlugin(p)
		p.SkipOSExit()
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "notifications command should not return an error:%s", err)
		assert.NotEmpty(t, out, "notifications command should not return an empty string")
		assert.Containsf(t, out, "1 notifications pending", "notifications command should one notification")
		assert.Containsf(t, out, "WARNING", "notifications command should raise warning at >= 1")
		t.Logf(out)
	})
	t.Run("notifications cmd with ignore", func(t *testing.T) {
		args := []string{
			"notifications",
			"--debug",
			"--warn", "0",
			"--ignore", "LOW_BAT",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		SetHmPlugin(p)
		p.SkipOSExit()
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "notifications command should not return an error:%s", err)
		assert.NotEmpty(t, out, "notifications command should not return an empty string")
		assert.Containsf(t, out, "0 notifications pending", "notifications command should not contain a notification")
		assert.Containsf(t, out, "ignoring notification: LOWBAT", "ignore should be mentioned")
		assert.Containsf(t, out, "OK", "notifications command should  have OK")
		t.Logf(out)
	})
}
