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

func TestSysvar(t *testing.T) {
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
	t.Run("sysvar list cmd", func(t *testing.T) {
		// reset map
		hmlib.SysVarIDMap = map[string]hmlib.SysVarEntry{}

		// mock the response for sysvar call
		response := SysVarListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := hmURL + hmlib.SysVarListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"sysvar",
			"list",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)

		// check result
		assert.NoErrorf(t, err, "sysvars command should not return an error:%s", err)
		assert.NotEmpty(t, out, "sysvars command should not return an empty string")
		assert.Containsf(t, out, "DutyCycle-LGW", "deviceList command should contain 'DutyCycle-LGW'")
		t.Logf(out)
	})

	t.Run("sysvar check cmd match value", func(t *testing.T) {
		// reset map
		hmlib.SysVarIDMap = map[string]hmlib.SysVarEntry{}

		// mock the response for sysvar call
		response := SysVarTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := hmURL + hmlib.SysVarEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"sysvar",
			"check",
			"--debug",
			"--unit-test",
			"--id", "8254",
			"--match", "xxx",
		}

		p := nagios.NewPlugin()
		SetPlugin(p)
		p.SkipOSExit()
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "sysvars command should not return an error:%s", err)
		assert.NotEmpty(t, out, "sysvars command should not return an empty string")
		assert.Containsf(t, out, "DutyCycle-LGW", "sysvars command should contain 'DutyCycle-LGW'")
		assert.Containsf(t, out, "CRITICAL", "DutyCycle command should raise CRITICAL if value not matched")
		_ = sysvarCheckCmd.Flags().Set("match", "")
		_ = sysvarCheckCmd.Flags().Set("id", "")
		t.Logf(out)
	})
	t.Run("sysvar check cmd warn value", func(t *testing.T) {
		// reset map
		hmlib.SysVarIDMap = map[string]hmlib.SysVarEntry{}

		// mock the response for sysvar call
		response := SysVarTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := hmURL + hmlib.SysVarEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"sysvar",
			"check",
			"--debug",
			"--unit-test",
			"--id", "8254",
			"--warn", "5",
		}

		p := nagios.NewPlugin()
		SetPlugin(p)
		p.SkipOSExit()
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "sysvars command should not return an error:%s", err)
		assert.NotEmpty(t, out, "sysvars command should not return an empty string")
		assert.Containsf(t, out, "DutyCycle-LGW", "sysvars command should contain 'DutyCycle-LGW'")
		assert.Containsf(t, out, "WARNING", "DutyCycle command should raise warning at >= 5")
		_ = sysvarCheckCmd.Flags().Set("warn", "")
		_ = sysvarCheckCmd.Flags().Set("id", "")
		t.Logf(out)
	})
}
