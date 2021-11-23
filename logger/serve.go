package logger

import "github.com/sirupsen/logrus"

func Info(args ...interface{}) {
	logger.Log(logrus.InfoLevel, args...)
}

func InfoF(format string, args ...interface{}) {
	logger.Logf(logrus.InfoLevel, format, args...)
}

func Warn(args ...interface{}) {
	logger.Log(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	logger.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	logger.Log(logrus.FatalLevel, args...)
	logger.Exit(1)
}

func Panic(args ...interface{}) {
	logger.Log(logrus.PanicLevel, args...)
}

func WarnF(format string, args ...interface{}) {
	logger.Logf(logrus.WarnLevel, format, args...)
}

func ErrorF(format string, args ...interface{}) {
	logger.Logf(logrus.ErrorLevel, format, args...)
}

func FatalF(format string, args ...interface{}) {
	logger.Logf(logrus.FatalLevel, format, args...)
	logger.Exit(1)
}

func PanicF(format string, args ...interface{}) {
	logger.Logf(logrus.PanicLevel, format, args...)
}
