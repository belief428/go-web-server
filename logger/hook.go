package logger

import (
	"encoding/json"
	"fmt"
	"github.com/belief428/go-web-server/utils"
	"os"
	"path/filepath"
	"runtime"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// Hook
type Hook struct {
	Path string
	File string
}

// Levels 只定义 error, warn, panic 等级的日志,其他日志等级不会触发 hook
func (this *Hook) Levels() []log.Level {
	return []log.Level{
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
	}
}

// Fire 将异常日志写入到指定日志文件中
func (this *Hook) Fire(entry *log.Entry) error {
	if isExist, _ := utils.PathExists(this.Path); !isExist {
		utils.MkdirAll(this.Path)
	}
	f, err := os.OpenFile(this.Path+this.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	_bytes, _ := json.Marshal(entry.Data)
	_, err = f.Write(_bytes)

	return err
}

func formatter(isConsole bool) *nested.Formatter {
	fmtter := &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	}
	if isConsole {
		fmtter.NoColors = false
	} else {
		fmtter.NoColors = true
	}
	return fmtter
}

func NewHook(logName string, rotationTime time.Duration, leastDay uint) log.Hook {
	writer, err := rotatelogs.New(
		// 日志文件
		logName+".%Y%m%d",
		rotatelogs.WithRotationCount(leastDay), // 只保留最近的N个日志文件
	)
	if err != nil {
		panic(err)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, formatter(false))

	return lfsHook
}
