package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var DebugFlag = false

// Init initializes the logger
func Init() {
	if DebugFlag {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

// Info logs an info message
func Info(msg string, fields ...interface{}) {
	if len(fields) == 0 {
		log.Info(msg)
	} else {
		log.WithFields(toFields(fields...)).Info(msg)
	}
}

// Error logs an error message
func Error(msg string, fields ...interface{}) {
	if len(fields) == 0 {
		log.Error(msg)
	} else {
		log.WithFields(toFields(fields...)).Error(msg)
	}
}

// Debug logs a debug message
func Debug(msg string, fields ...interface{}) {
	if len(fields) == 0 {
		log.Debug(msg)
	} else {
		log.WithFields(toFields(fields...)).Debug(msg)
	}
}

// Warn logs a warning message
func Warn(msg string, fields ...interface{}) {
	if len(fields) == 0 {
		log.Warn(msg)
	} else {
		log.WithFields(toFields(fields...)).Warn(msg)
	}
}

// toFields converts a slice of interface{} to logrus.Fields
func toFields(fields ...interface{}) logrus.Fields {
	if len(fields)%2 != 0 {
		// If odd number of fields, last one is treated as a value with key "value"
		fields = append(fields, "value")
	}

	result := make(logrus.Fields)
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			key = "unknown"
		}
		result[key] = fields[i+1]
	}
	return result
}