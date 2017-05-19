// Package logger provide a very simple logger with levels
//
// Example:
//
//   import "logger"
//
//   log := GetLogger(INFO, "logger")
//
//   log.Info("should print")
//
//   log.Debug("current level is %v", log.Level) // but this will not print
//
//   log.Fatal("print and die!")
//
package logger

import (
	"log"
	"os"
	"strings"
)

// LogLevel defines the different log levels
type LogLevel int

const (
	ERROR   LogLevel = -1
	WARNING LogLevel = 0
	INFO    LogLevel = 1
	DEBUG   LogLevel = 2
)

var labels = []string{
	"ERROR   ",
	"WARNING ",
	"INFO    ",
	"DEBUG   ",
}

// Logger represent an active logging object, with the current level and system logger
type Logger struct {
	Level  LogLevel    // current log level
	Logger *log.Logger // system logger
}

// GetLogger returns an instance of a logger, with LogLevel set to the specified level.
//
// The system logger is initialized with the specified prefix and the standard system logger defaults
// (write to os.Stderr in the default log format)
func GetLogger(level LogLevel, prefix string) *Logger {
	if len(prefix) > 0 && !strings.HasSuffix(prefix, " ") {
		prefix += " "
	}
	return &Logger{Level: level, Logger: log.New(os.Stderr, prefix, log.LstdFlags)}
}

// SetLevel set the new level for this logger
func (l *Logger) SetLevel(level LogLevel) {
	l.Level = level
}

// Log logs the specified message, with arguments if Logger.LogLevel >= level
func (l *Logger) Log(level LogLevel, fmt string, args ...interface{}) {
	if level < 0 {
		l.Logger.Printf(labels[1+ERROR]+fmt, args...)
	} else if l.Level >= level {
		ilevel := int(level)

		if ilevel > len(labels)-1 {
			ilevel = len(labels) - 1
		}
		l.Logger.Printf(labels[1+ilevel]+fmt, args...)
	}
}

// Debug logs the specified message if this Logger.LogLevel is DEBUG or lower
func (l *Logger) Debug(fmt string, args ...interface{}) {
	l.Log(DEBUG, fmt, args...)
}

// Info logs the specified message if this Logger.LogLevel is INFO or lower
func (l *Logger) Info(fmt string, args ...interface{}) {
	l.Log(INFO, fmt, args...)
}

// Warning logs the specified message if this Logger.LogLevel is WARNING or lower
func (l *Logger) Warning(fmt string, args ...interface{}) {
	l.Log(WARNING, fmt, args...)
}

// Error always logs the specified message
func (l *Logger) Error(fmt string, args ...interface{}) {
	l.Log(ERROR, fmt, args...)
}

// Fatal always logs the specified message (at ERROR level) and terminate the application
func (l *Logger) Fatal(fmt string, args ...interface{}) {
	l.Log(ERROR, fmt, args...)
	os.Exit(1)
}
