package cmd

import (
	"testing"

	"github.com/atc0005/go-nagios"
	"github.com/stretchr/testify/assert"
)

func TestNagios(t *testing.T) {
	// test table
	tests := []struct {
		name   string
		status string
		// https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT
		perfdata       []nagios.PerformanceData
		expectedStatus int
	}{
		{"nagiosOK", "OK", nil, nagios.StateOKExitCode},
		{"nagiosWarn", "WARNING", nil, nagios.StateWARNINGExitCode},
		{"nagiosCrtit", "CRITICAL", nil, nagios.StateCRITICALExitCode},
		{"nagiosUnknown", "UNKNOWN", nil, nagios.StateUNKNOWNExitCode},
		{"nagiosNoStatus", "", nil, nagios.StateUNKNOWNExitCode},
		{"nagiosPerfCritical", "",
			[]nagios.PerformanceData{{Label: "Test", Value: "100", UnitOfMeasurement: "", Warn: "~:9", Crit: "99", Min: "0", Max: "1000"}},
			nagios.StateCRITICALExitCode},
		{"nagiosPerfWarning", "",
			[]nagios.PerformanceData{{Label: "Test", Value: "3", Warn: "1", Crit: "10"}},
			nagios.StateWARNINGExitCode},
	}

	for _, tt := range tests {
		checkHM := nagios.NewPlugin()
		SetPlugin(checkHM)
		checkHM.SkipOSExit()
		t.Run(tt.name, func(t *testing.T) {
			p := NagiosResult(tt.status, "no notifications", "", tt.perfdata)
			sc := p.ExitStatusCode
			so := p.ServiceOutput
			lo := p.LongServiceOutput
			t.Logf("Output: %s\nLongOutput: %s", so, lo)
			assert.Equal(t, tt.expectedStatus, sc, "NagiosResult should return %s", tt.status)
			assert.Zerof(t, len(p.Errors), "result should not have errors, but have thise: %v", p.Errors)
		})
	}
}
