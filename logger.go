package lord

import (
	"errors"
)

var log Logger

// Fields ...
type Fields map[string]interface{}

const (
	// Debug string
	Debug = "debug"
	// Info string
	Info = "info"
	// Warn string
	Warn = "warning"
	// Error string
	Error = "error"
	// Fatal string
	Fatal = "fatal"
)

const (
	// InstanceZapLogger iota 0
	InstanceZapLogger int = iota
	// InstanceLogrusLogger iota 1
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

// Logger interface ...
type Logger interface {
	Debugf(f string, args ...interface{})
	Infof(f string, args ...interface{})
	Warnf(f string, args ...interface{})
	Errorf(f string, args ...interface{})
	Fatalf(f string, args ...interface{})
	Panicf(f string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

// Configuration ...
type Configuration struct {
	EnableConsole     bool   `mapstructure:"enableConsole"`
	ConsoleJSONFormat bool   `mapstructure:"consoleJsonFormat"`
	ConsoleLevel      string `mapstructure:"consoleLevel"`
	EnableFile        bool   `mapstructure:"enableFile"`
	FileJSONFormat    bool   `mapstructure:"fileJsonFormat"`
	FileLevel         string `mapstructure:"fileLevel"`
	FileLocation      string `mapstructure:"fileLocation"`
}

//NewLogger returns an instance of logger
func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil
	default:
		return errInvalidLoggerInstance
	}
}

// Debugf formating string
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof formating string
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf formating string
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf formating string
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf formating string
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf formating string
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields ...
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
