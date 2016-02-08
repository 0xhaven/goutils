package log

import (
	"flag"
	"log"
	"log/syslog"
	"os"
)

const (
	// LevelDebug is the log level for Debug statements.
	LevelDebug = iota
	// LevelInfo is the log level for Info statements.
	LevelInfo
	// LevelWarning is the log level for Warning statements.
	LevelWarning
	// LevelError is the log level for Error statements.
	LevelError
	// LevelCritical is the log level for Critical statements.
	LevelCritical
	// LevelFatal is the log level for Fatal statements.
	LevelFatal
)

var levelPrefix = [...]string{
	LevelDebug:    "[DEBUG] ",
	LevelInfo:     "[INFO] ",
	LevelWarning:  "[WARNING] ",
	LevelError:    "[ERROR] ",
	LevelCritical: "[CRITICAL] ",
	LevelFatal:    "[FATAL] ",
}

var levelPriority = [...]syslog.Priority{
	LevelDebug:    syslog.LOG_DEBUG,
	LevelInfo:     syslog.LOG_INFO,
	LevelWarning:  syslog.LOG_WARNING,
	LevelError:    syslog.LOG_ERR,
	LevelCritical: syslog.LOG_CRIT,
	LevelFatal:    syslog.LOG_EMERG,
}

var (
	// Level is the minimum logging level to log at.
	Level int
	// Loggers maps each logging level to a *log.Logger that will be used for it.
	Loggers = map[int]*log.Logger{}
)

func init() {
	var useSyslog bool
	flag.IntVar(&Level, "loglevel", LevelInfo, "Minimum log level")
	flag.BoolVar(&useSyslog, "syslog", false, "Whether or not to use syslog logging")
	flag.Parse()

	for l := LevelDebug; l <= LevelFatal; l++ {
		if !useSyslog {
			Loggers[l] = log.New(os.Stderr, levelPrefix[l], 0)
		} else {
			var err error
			Loggers[l], err = syslog.NewLogger(levelPriority[l], 0)
			if err != nil {
				log.Fatal(err)
			}

			Loggers[l].SetPrefix(levelPrefix[l])
		}
	}
}

func outputf(l int, format string, v []interface{}) {
	if l >= Level {
		Loggers[l].Printf(format, v...)
	}
}

func output(l int, v []interface{}) {
	if l >= Level {
		Loggers[l].Print(v...)
	}
}

// Fatalf logs a formatted message at the "fatal" level and then exits. The
// arguments are handled in the same manner as fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	outputf(LevelFatal, format, v)
	os.Exit(1)
}

// Fatal logs its arguments at the "fatal" level and then exits.
func Fatal(v ...interface{}) {
	output(LevelFatal, v)
	os.Exit(1)
}

// Criticalf logs a formatted message at the "critical" level. The
// arguments are handled in the same manner as fmt.Printf.
func Criticalf(format string, v ...interface{}) {
	outputf(LevelCritical, format, v)
}

// Critical logs its arguments at the "critical" level.
func Critical(v ...interface{}) {
	output(LevelCritical, v)
}

// Errorf logs a formatted message at the "error" level. The arguments
// are handled in the same manner as fmt.Printf.
func Errorf(format string, v ...interface{}) {
	outputf(LevelError, format, v)
}

// Error logs its arguments at the "error" level.
func Error(v ...interface{}) {
	output(LevelError, v)
}

// Warningf logs a formatted message at the "warning" level. The
// arguments are handled in the same manner as fmt.Printf.
func Warningf(format string, v ...interface{}) {
	outputf(LevelWarning, format, v)
}

// Warning logs its arguments at the "warning" level.
func Warning(v ...interface{}) {
	output(LevelWarning, v)
}

// Infof logs a formatted message at the "info" level. The arguments
// are handled in the same manner as fmt.Printf.
func Infof(format string, v ...interface{}) {
	outputf(LevelInfo, format, v)
}

// Info logs its arguments at the "info" level.
func Info(v ...interface{}) {
	output(LevelInfo, v)
}

// Debugf logs a formatted message at the "debug" level. The arguments
// are handled in the same manner as fmt.Printf.
func Debugf(format string, v ...interface{}) {
	outputf(LevelDebug, format, v)
}

// Debug logs its arguments at the "debug" level.
func Debug(v ...interface{}) {
	output(LevelDebug, v)
}
