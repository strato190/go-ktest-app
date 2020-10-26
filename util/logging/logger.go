package logging

import (
	log "github.com/sirupsen/logrus"
)

const (
	app = "go-ktest-app"
)

var (
	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// Setup initialize the log instance
func Setup() {
	logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.SetLevel(log.DebugLevel)
	printEvent(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.SetLevel(log.InfoLevel)
	printEvent(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.SetLevel(log.WarnLevel)
	printEvent(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.SetLevel(log.ErrorLevel)
	printEvent(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.SetLevel(log.FatalLevel)
	printEvent(v)
}

// printEvent set the prefix of the log output
func printEvent(v ...interface{}) {
	logger.WithFields(log.Fields{
		"app": app,
	}).Println(v)
}
