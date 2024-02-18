package cmd

import (
	"regexp"

	"github.com/atc0005/go-nagios"
	log "github.com/sirupsen/logrus"
)

var hmPlugin = nagios.NewPlugin()

const (
	stateOK       = "OK"
	stateWarning  = "WARNING"
	stateCritical = "CRITICAL"
	stateUnknown  = "UNKNOWN"
)

// NagiosResult is a helper function to return a nagios plugin result
func NagiosResult(status string, output string, longOutput string, perfdata []nagios.PerformanceData) *nagios.Plugin {
	p := GetHmPlugin()

	// Second, immediately defer ReturnCheckResults() so that it runs as the
	// last step in your client code. If you do not defer ReturnCheckResults()
	// immediately any other deferred functions in your client code will not
	// run.
	//
	// Avoid calling os.Exit() directly from your code. If you do, this
	// library is unable to function properly; this library expects that it
	// will handle calling os.Exit() with the required exit code (and
	// specifically formatted output).
	//
	// For handling error cases, the approach is roughly the same, only you
	// call return explicitly to end execution of the client code and allow
	// deferred functions to run.
	defer p.ReturnCheckResults()

	// more stuff here

	p.ServiceOutput = output
	p.LongServiceOutput = longOutput
	switch status {
	case stateOK:
		p.ExitStatusCode = nagios.StateOKExitCode
	case stateWarning:
		p.ExitStatusCode = nagios.StateWARNINGExitCode
	case stateCritical:
		p.ExitStatusCode = nagios.StateCRITICALExitCode
	default:
		p.ExitStatusCode = nagios.StateUNKNOWNExitCode
	}
	if len(hmWarnThreshold) > 0 {
		p.WarningThreshold = hmWarnThreshold
	}
	if len(hmCritThreshold) > 0 {
		p.CriticalThreshold = hmCritThreshold
	}
	if len(perfdata) > 0 {
		err := p.AddPerfData(false, perfdata...)
		if err != nil {
			p.Errors = append(p.Errors, err)
		}
		err = p.EvaluateThreshold(perfdata...)
		if err != nil {
			p.Errors = append(p.Errors, err)
		}
	}
	sl := nagios.SupportedStateLabels()
	status = sl[p.ExitStatusCode]
	log.Debugf("Result: %s:%s", status, output)
	return p
}

// SetHmPlugin sets the nagios plugin for the library
func SetHmPlugin(plugin *nagios.Plugin) {
	hmPlugin = plugin
}

// GetHmPlugin returns the nagios plugin for the library
func GetHmPlugin() *nagios.Plugin {
	return hmPlugin
}

func isNumeric(s string) bool {
	re := regexp.MustCompile(`^[+-\\d.]+$`)
	r := re.MatchString(s)
	return r
}
