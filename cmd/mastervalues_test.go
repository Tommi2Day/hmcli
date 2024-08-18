package cmd

import (
	"testing"

	"github.com/tommi2day/hmcli/test"

	"github.com/atc0005/go-nagios"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

func TestMasterValues(t *testing.T) {
	var httpClient = resty.New()
	test.InitTestDirs()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	defer httpmock.DeactivateAndReset()
	hmlib.SetHTTPClient(httpClient)
	hmToken = MockToken
	hmURL = MockURL
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)

	response := MasterValueTest
	responder := httpmock.NewStringResponder(200, response)
	fakeURL := MockURL + hmlib.MasterValueEndpoint
	httpmock.RegisterResponder("GET", fakeURL, responder)

	response = DeviceListTest
	responder = httpmock.NewStringResponder(200, response)
	fakeURL = MockURL + hmlib.DeviceListEndpoint
	httpmock.RegisterResponder("GET", fakeURL, responder)
	defer httpmock.DeactivateAndReset()

	t.Run("MasterValue list NoId", func(t *testing.T) {
		args := []string{
			"mastervalues",
			"list",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.Errorf(t, err, "mastervalues command should return an error")
		assert.Containsf(t, out, "please provide an address or id", "Id should missing")
		t.Log(out)
	})
	t.Run("MasterValue list", func(t *testing.T) {
		args := []string{
			"mastervalues",
			"list",
			"--id", "4740, 4763",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "mastervalues command should not return an error:%s", err)
		assert.NotEmpty(t, out, "mastervalues command should not return an empty string")
		assert.Containsf(t, out, "MOTIONDETECTOR_TRANSCEIVER", "mastervalues command should contain MOTIONDETECTOR_TRANSCEIVER")
		_ = valueListCmd.Flags().Set("id", "")
		t.Log(out)
	})
	t.Run("MasterValue Check", func(t *testing.T) {
		args := []string{
			"mastervalues",
			"check",
			"--id", "4763",
			"--name", "MIN_INTERVAL",
			"--warn", "1",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "mastervalues command should not return an error:%s", err)
		assert.NotEmpty(t, out, "mastervalues command should not return an empty string")
		assert.Containsf(t, out, "MOTIONDETECTOR_TRANSCEIVER", "mastervalues command should contain MOTIONDETECTOR_TRANSCEIVER")
		assert.Containsf(t, out, "WARNING", "mastervalues command should raise warning at >= 1")
		// reset warn and id flag
		_ = valueCheckCmd.Flags().Set("warn", "")
		_ = valueCheckCmd.Flags().Set("id", "")
		t.Log(out)
	})
	t.Run("MasterValue match Check", func(t *testing.T) {
		args := []string{
			"mastervalues",
			"check",
			"--id", "4763",
			"--name", "MIN_INTERVAL",
			"--match", "[^\\d+]",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "mastervalues command should not return an error:%s", err)
		assert.NotEmpty(t, out, "mastervalues command should not return an empty string")
		assert.Containsf(t, out, "MOTIONDETECTOR_TRANSCEIVER", "mastervalues command should contain MOTIONDETECTOR_TRANSCEIVER")
		assert.Containsf(t, out, "CRITICAL", "mastervalues command should raise Critical if not matched")
		// reset match and id flag
		_ = valueCheckCmd.Flags().Set("match", "")
		_ = valueCheckCmd.Flags().Set("id", "")
		t.Log(out)
	})
	t.Run("MasterValue wrong id Check", func(t *testing.T) {
		args := []string{
			"mastervalues",
			"check",
			"--id", "4743",
			"--name", "MIN_INTERVAL",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.Errorf(t, err, "mastervalues command should not return an error:%s", err)
		assert.Containsf(t, out, "device with id 4743 not in response",
			"mastervalues command should contain 'device with id 4743 not in response'")
		// reset warn flag
		_ = valueCheckCmd.Flags().Set("name", "")
		t.Log(out)
	})
}
