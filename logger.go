package gocbadaptor

import "log"

// SilentLogger logger say noting
type SilentLogger struct{}

// Log but say nothing
func (s *SilentLogger) Log(v ...interface{}) {}

// Log but say nothing with format either
func (s *SilentLogger) Logf(v ...interface{}) {}

// DefaultLogger type-default logger
type DefaultLogger struct {
	enabled bool
}

// NewDefaultLogger Generate New Logger
func NewDefaultLogger(enabled bool) Loggerable {
	return &DefaultLogger{
		enabled: enabled,
	}
}

// LogEnabled return Log-enabled or not
func (d *DefaultLogger) LogEnabled() bool {
	return d.enabled
}

// Log logging with log-package
func (d *DefaultLogger) Log(v ...interface{}) {
	if d.LogEnabled() {
		log.Println(v...)
	}
}

// Logf logging with format using log-package
func (d *DefaultLogger) Logf(f string, v ...interface{}) {
	if d.LogEnabled() {
		log.Printf(f, v...)
	}
}
