package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var stdoutLogger = logrus.New()
var stderrLogger = logrus.New()

func init() {
	Setup()
}
func Setup() {
	stderrLogger.SetOutput(os.Stderr)
	stderrLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})
	stdoutLogger.SetOutput(os.Stdout)
	stdoutLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})
}

const (
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
)

func Log(level logrus.Level, args ...interface{}) {
	if level >= ErrorLevel {
		stderrLogger.Log(level, args...)
	} else {
		stdoutLogger.Log(level, args...)
	}
}
func Error(args ...interface{}) {
	Log(ErrorLevel, args...)
}
func Warn(args ...interface{}) {
	Log(WarnLevel, args...)
}
func Info(args ...interface{}) {
	Log(InfoLevel, args...)
}
func Debug(args ...interface{}) {
	Log(DebugLevel, args...)
}
func Logf(level logrus.Level, format string, args ...interface{}) {
	if level >= ErrorLevel {
		stderrLogger.Logf(level, format, args...)
	}
}
func Errorf(format string, args ...interface{}) {
	Logf(ErrorLevel, format, args...)
}
func Warnf(format string, args ...interface{}) {
	Logf(WarnLevel, format, args...)
}
func Infof(format string, args ...interface{}) {
	Logf(InfoLevel, format, args...)
}
func Debugf(format string, args ...interface{}) {
	Logf(DebugLevel, format, args...)
}
func Printf(format string, args ...interface{}) {
	Infof(format, args...)
}
func Println(args ...interface{}) {
	Info(args...)
}
func Fatal(args ...interface{}) {
	Error(args...)
	os.Exit(1)
}
func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(args ...interface{}) {
	Error(args...)
	Errorf("\n")
	os.Exit(1)
}
