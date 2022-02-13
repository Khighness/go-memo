package util

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// @Author KHighness
// @Update 2022-02-13

var logger *logrus.Logger

const (
	DateFormat string = "2006-01-02"
	TimeFormat string = "2006-01-02 15:04:05"
	FileSuffix string = ".logger"
)

func Logger() *logrus.Logger {
	src, _ := logfile()

	// 已实例化
	if logger != nil {
		logger.Out = src
		return logger
	}

	newLogger := logrus.New()
	newLogger.Out = src
	newLogger.SetLevel(logrus.DebugLevel)
	newLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		TimestampFormat:           TimeFormat,
	})
	logger = newLogger
	return logger
}

// 设置日志输出文件
func logfile() (*os.File, error) {
	// 根据日期生成文件名
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format(DateFormat) + FileSuffix
	logFile := path.Join(logFilePath, logFileName)

	// 创建日志文件
	if _, err := os.Stat(logFile); err != nil {
		if _, err := os.Create(logFile); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	// 写入文件
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return src, nil
}