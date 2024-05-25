package cmd

import (
	"github.com/tommi2day/check_hm/test"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

func TestDevices(t *testing.T) {
	var httpClient = resty.New()
	test.Testinit(t)
	httpmock.Reset()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	hmlib.SetHTTPClient(httpClient)
	defer httpmock.DeactivateAndReset()
	hmToken = MockToken
	hmURL = MockURL
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)

	t.Run("deviceList cmd", func(t *testing.T) {
		response := DeviceListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.DeviceListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"device",
			"list",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "deviceList command should not return an error:%s", err)
		assert.NotEmpty(t, out, "deviceList command should not return an empty string")
		assert.Containsf(t, out, "Bewegungsmelder Garage", "deviceList command should contain Bewegungsmelder Garage")
		t.Logf(out)
	})
	t.Run("deviceList cmd by name", func(t *testing.T) {
		response := DeviceListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.DeviceListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"device",
			"list",
			"--name", "Bewegungsmelder Garage",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "deviceList command should not return an error:%s", err)
		assert.NotEmpty(t, out, "deviceList command should not return an empty string")
		assert.Contains(t, out, "Bewegungsmelder Garage", "deviceList command should contain Bewegungsmelder Garage")
		assert.Contains(t, out, "found", "deviceList command should contain 'found'")
		_ = deviceListCmd.Flags().Set("name", "")
		t.Logf(out)
	})
	t.Run("deviceList cmd by ID", func(t *testing.T) {
		response := DeviceListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.DeviceListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"device",
			"list",
			"--id", "4740",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "deviceList command should not return an error:%s", err)
		assert.NotEmpty(t, out, "deviceList command should not return an empty string")
		assert.Containsf(t, out, "Bewegungsmelder Garage", "deviceList command should contain Bewegungsmelder Garage")
		assert.Contains(t, out, "found", "deviceList command should contain 'found'")
		_ = deviceListCmd.Flags().Set("id", "")
		t.Logf(out)
	})
	t.Run("deviceList cmd by address", func(t *testing.T) {
		response := DeviceListTest
		responder := httpmock.NewStringResponder(200, response)
		fakeURL := MockURL + hmlib.DeviceListEndpoint
		httpmock.RegisterResponder("GET", fakeURL, responder)
		args := []string{
			"device",
			"list",
			"--address", "000955699D3D84",
			"--debug",
			"--unit-test",
		}
		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "deviceList command should not return an error:%s", err)
		assert.NotEmpty(t, out, "deviceList command should not return an empty string")
		assert.Containsf(t, out, "Bewegungsmelder Garage", "deviceList command should contain Bewegungsmelder Garage")
		assert.Contains(t, out, "found", "deviceList command should contain 'found'")
		_ = deviceListCmd.Flags().Set("address", "")
		t.Logf(out)
	})
}
