package log

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

var (
	logger      = log.Default()
	errorLogger = log.New(os.Stderr, "", 0)
	level       = LevelDefault
	name        string
)

func init() {
	logger.SetFlags(0)
}

func Init(n string, severity Level) {
	level = severity
	name = n
}

func printIfLevel(log func(v ...interface{}), severity Level, message string, v ...interface{}) {
	if level <= severity {
		log(&Entry{
			Message:   fmt.Sprintf(message, v...),
			Severity:  severity.String(),
			Trace:     "",
			Component: name,
		})
	}
}

func logIfLevel(severity Level, message string, v ...interface{}) {
	printIfLevel(logger.Print, severity, message, v...)
}

func errorIfLevel(severity Level, message string, v ...interface{}) {
	printIfLevel(errorLogger.Print, severity, message, v...)
}

func errorWithStackIfLevel(severity Level, message string, v ...interface{}) {
	errorIfLevel(severity, message+":\n%s", append(v, debug.Stack())...)
}

func DefaultLogger() *log.Logger {
	return logger
}

func Debug(format string, v ...interface{}) {
	logIfLevel(LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	logIfLevel(LevelInfo, format, v...)
}

func Notice(format string, v ...interface{}) {
	logIfLevel(LevelNotice, format, v...)
}

func Warning(format string, v ...interface{}) {
	logIfLevel(LevelWarning, format, v...)
}

func Error(format string, v ...interface{}) {
	logIfLevel(LevelError, format, v...)
}

func Critical(format string, v ...interface{}) {
	errorWithStackIfLevel(LevelCritical, format, v...)
}

func Alert(format string, v ...interface{}) {
	errorWithStackIfLevel(LevelAlert, format, v...)
}

func Emergency(format string, v ...interface{}) {
	errorWithStackIfLevel(LevelEmergency, format, v...)
}
