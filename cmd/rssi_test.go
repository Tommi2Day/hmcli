package cmd

import (
	"testing"

	"github.com/tommi2day/hmcli/test"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tommi2day/gomodules/common"
	"github.com/tommi2day/gomodules/hmlib"
)

func TestRssi(t *testing.T) {
	var httpClient = resty.New()
	test.InitTestDirs()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	defer httpmock.DeactivateAndReset()
	hmlib.SetHTTPClient(httpClient)
	hmToken = MockToken
	hmURL = MockURL
	hmlib.SetHmToken(hmToken)
	hmlib.SetHmURL(hmURL)

	response := RssiTest
	responder := httpmock.NewStringResponder(200, response)
	fakeURL := MockURL + hmlib.RssiEndpoint
	httpmock.RegisterResponder("GET", fakeURL, responder)
	defer httpmock.DeactivateAndReset()

	t.Run("Rssi cmd", func(t *testing.T) {
		args := []string{
			"rssi",
			"--debug",
			"--unit-test",
		}

		out, err := common.CmdRun(RootCmd, args)
		assert.NoErrorf(t, err, "rssi command should not return an error:%s", err)
		assert.NotEmpty(t, out, "rssi command should not return an empty string")
		assert.Containsf(t, out, "MEQ0481419", "rssi command should contain MEQ0481419")
		t.Logf(out)
	})
}
