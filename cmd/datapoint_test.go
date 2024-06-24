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

func TestDatapoint(t *testing.T) {
	var httpClient = resty.New()
	test.InitTestDirs()
	httpmock.Reset()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	hmlib.SetHTTPClient(httpClient)
	defer httpmock.DeactivateAndReset()
	hmToken = MockToken
	hmURL = MockURL
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)

	t.Run("datapoint list cmd", func(t *testing.T) {
		response := StateListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.StateListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"datapoint",
			"list",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "datapoint List command should not return an error:%s", err)
		assert.NotEmpty(t, out, "datapoint List command should not return an empty string")
		assert.Contains(t, out, "SUCCESS: 33 data points", "deviceList command should contain 'SUCCESS: 33 data points'")
		t.Logf(out)
	})
	t.Run("datapoint list cmd with match", func(t *testing.T) {
		response := StateListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.StateListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"datapoint",
			"list",
			"--match", "CONFIG_PENDING",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "datapoint List command should not return an error:%s", err)
		assert.NotEmpty(t, out, "datapoint List command should not return an empty string")
		assert.Contains(t, out, "SUCCESS: 2 data points", "deviceList command should contain 'SUCCESS: 2 data points'")
		_ = datapointListCmd.Flags().Set("match", "")
		t.Logf(out)
	})

	t.Run("datapoint Check", func(t *testing.T) {
		args := []string{
			"datapoint",
			"check",
			"--id", "7799",
			"--warn", "100",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "datapoint check command should not return an error:%s", err)
		assert.NotEmpty(t, out, "datapoint check command should not return an empty string")
		assert.Containsf(t, out, "HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION", "datapoint command should contain HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION")
		assert.Containsf(t, out, "WARNING", "datapoint command should raise warning at >= 1")
		// reset warn flag and id
		_ = datapointCheckCmd.Flags().Set("warn", "")
		_ = datapointCheckCmd.Flags().Set("id", "")
		t.Logf(out)
	})
	t.Run("datapoint per name match Check", func(t *testing.T) {
		args := []string{
			"datapoint",
			"check",
			"--name", "HmIP-RF.000955699D3D84:0.CONFIG_PENDING",
			"--match", "true",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "datapoint command should not return an error:%s", err)
		assert.NotEmpty(t, out, "datapoint command should not return an empty string")
		assert.Containsf(t, out, "", "mastervalues command should contain MOTIONDETECTOR_TRANSCEIVER")
		assert.Containsf(t, out, "CRITICAL", "mastervalues command should raise Critical if not matched")
		// reset name and match flag
		_ = datapointCheckCmd.Flags().Set("match", "")
		_ = datapointCheckCmd.Flags().Set("name", "")
		t.Logf(out)
	})
	t.Run("datapoint wrong id Check", func(t *testing.T) {
		args := []string{
			"datapoint",
			"check",
			"--id", "4743",
			"--debug",
			"--unit-test",
		}
		p := nagios.NewPlugin()
		p.SkipOSExit()
		SetPlugin(p)
		out, err := common.CmdRun(RootCmd, args)
		assert.Errorf(t, err, "datapoint command should not return an error:%s", err)
		assert.Containsf(t, out, "UNKNOWN: datapoint id 4743 not found",
			"datapoint command should contain 'UNKNOWN: datapoint id 4743 not found'")
		// reset warn flag
		_ = datapointCheckCmd.Flags().Set("id", "")
		t.Logf(out)
	})
}
