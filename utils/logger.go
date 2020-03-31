/*
 * Package utils
 * File: logger.go
 * Project: SXYGo
 * File Created: 2020-03-29 17:55:16
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-31 18:43:27
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package utils

import (
	"os"

	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var logger *Logger

// Logger 日志
type Logger struct {
	level  logrus.Level
	logrus *logrus.Logger
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	intLevel := logrus.ErrorLevel
	switch level {
	case "error":
		intLevel = logrus.ErrorLevel
	case "warning":
		intLevel = logrus.WarnLevel
	case "info":
		intLevel = logrus.InfoLevel
	case "debug":
		intLevel = logrus.DebugLevel
	}
	ll := logrus.New()
	ll.SetOutput(os.Stdout)
	ll.SetLevel(intLevel)
	ll.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "req", "ip"},
		TimestampFormat: "[2006-01-02 15:04:05]",
	})
	ll.AddHook(newLfsHook())

	l := Logger{
		level:  intLevel,
		logrus: ll,
	}
	logger = &l
}

// Log 返回日志对象
func Log() *logrus.Logger {
	if logger == nil {
		BuildLogger("debug")
	}
	return logger.logrus
}

// RequestLog 返回请求日志对象
func RequestLog(c *gin.Context) *logrus.Entry {
	if logger == nil {
		BuildLogger("debug")
	}
	ll := logger.logrus
	reqFields := c.MustGet("RequestFields")
	return ll.WithFields(reqFields.(logrus.Fields))
}

func newLfsHook() logrus.Hook {
	logName := "logs/sxygo"
	writer, err := rotatelogs.New(
		logName+"-%Y%m%d%H.log",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName("logs/latest.log"),

		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(14),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "[2006-01-02 15:04:05]",
		NoColors:        true,
	})

	return lfsHook
}
