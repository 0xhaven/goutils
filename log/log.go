package log

import (
	"flag"
	"fmt"
	"io"
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
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarning:  "WARNING",
	LevelError:    "ERROR",
	LevelCritical: "CRITICAL",
	LevelFatal:    "FATAL",
}

func newLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, fmt.Sprintf("[%s] ", prefix), 0)
}

// Loggers maps each logging level to a *log.Logger that will be used for it.
var Loggers = [...]*log.Logger{
	LevelDebug:    newLogger(os.Stderr, levelPrefix[LevelDebug]),
	LevelInfo:     newLogger(os.Stderr, levelPrefix[LevelInfo]),
	LevelWarning:  newLogger(os.Stderr, levelPrefix[LevelWarning]),
	LevelError:    newLogger(os.Stderr, levelPrefix[LevelError]),
	LevelCritical: newLogger(os.Stderr, levelPrefix[LevelCritical]),
	LevelFatal:    newLogger(os.Stderr, levelPrefix[LevelFatal]),
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
	// SysLoggers maps each logging level to a *log.Logger that will be used for it.
	SysLoggers = make([]*log.Logger, LevelFatal+1)
	// Syslog determines whether or not to log to syslog
	Syslog bool
	// SyslogTag is the tag to use for syslog
	SyslogTag string
	// SyslogNetwork is the remote syslog network.
	SyslogNetwork string
	// SyslogRemote is the remote syslog addr.
	SyslogRemote string
)

func init() {
	flag.IntVar(&Level, "loglevel", LevelInfo, "Minimum log level")
	flag.BoolVar(&Syslog, "syslog", false, "Whether or not to log to syslog")
	flag.StringVar(&SyslogTag, "syslog-tag", "", "Syslog tag to use")
	flag.StringVar(&SyslogNetwork, "syslog-network", "", "Syslog network to use")
	flag.StringVar(&SyslogRemote, "syslog-remote", "", "Syslog server to use (Defaults to localhost)")
}

func logger(l int) *log.Logger {
	if Syslog {
		if SysLoggers[l] != nil {
			return SysLoggers[l]
		}

		w, err := syslog.Dial(SyslogNetwork, SyslogRemote, levelPriority[l], SyslogTag)
		if err == nil {
			SysLoggers[l] = newLogger(w, levelPrefix[l])
			return SysLoggers[l]
		}

		Loggers[LevelError].Println("Unable to dial syslog:", err)
	}

	return Loggers[l]
}

func output(l int, v []interface{}) {
	if l >= Level {
		logger(l).Print(v...)
	}
}

func outputf(l int, format string, v []interface{}) {
	if l >= Level {
		logger(l).Printf(format, v...)
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
