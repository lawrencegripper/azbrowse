package tfprovider

import (
	"github.com/go-logr/logr"
)

type NullLogger struct {
}

var _ logr.Logger = &NullLogger{}

// Info implements Logger.Info
func (l *NullLogger) Info(msg string, kvs ...interface{}) {
}

// Enabled implements Logger.Enabled
func (*NullLogger) Enabled() bool {
	return false
}

// Error implements Logger.Error
func (l *NullLogger) Error(err error, msg string, kvs ...interface{}) {
	kvs = append(kvs, "error", err)
	l.Info(msg, kvs...)
}

// V implements Logger.V
func (l *NullLogger) V(_ int) logr.Logger {
	return l
}

// WithName implements Logger.WithName
func (l *NullLogger) WithName(name string) logr.Logger {
	return l
}

// WithValues implements Logger.WithValues
func (l *NullLogger) WithValues(kvs ...interface{}) logr.Logger {
	return l
}

// NewNullLogger creates a new NullLogger
func NewNullLogger() logr.Logger {
	return &NullLogger{}
}
