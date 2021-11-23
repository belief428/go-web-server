package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*Option
	Logger *log.Logger
}

type Option struct {
	File     string `json:"file"`
	LeastDay uint   `json:"least_day"`
	Level    string `json:"level"`
	IsStdout bool   `json:"is_stdout"`
}

var logger *log.Logger

var loggerLevel = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
}

func (this *Logger) level() log.Level {
	if _, has := loggerLevel[this.Level]; !has {
		return log.ErrorLevel
	}
	return loggerLevel[this.Level]
}

func (this *Logger) Load() {
	logger = this.Logger
}

func (this *Logger) Init(option *Option) *Logger {
	this.Option = option

	_logger := log.New()
	_logger.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	_logger.SetReportCaller(true)
	_logger.AddHook(NewHook(this.File, 0, this.LeastDay))

	if this.IsStdout {
		_logger.SetOutput(io.MultiWriter(os.Stdout))
	}
	_logger.SetFormatter(formatter(true))
	_logger.SetLevel(this.level())

	this.Logger = _logger

	return this
}

func NewLogger() *Logger {
	return &Logger{}
}
