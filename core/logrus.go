package core

import (
	"bytes"
	"fmt"
	"gvb_server/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n",
			global.Config.Logger.Prefix,
			timestamp,
			levelColor,
			entry.Level,
			fileVal,
			funcVal,
			entry.Message,
		)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n",
			global.Config.Logger.Prefix,
			timestamp,
			levelColor,
			entry.Level,
			entry.Message,
		)
	}
	return b.Bytes(), nil
}

func Logger() *logrus.Logger {
	mLog := logrus.New()
	//输出类型
	mLog.SetOutput(os.Stdout)
	//开启返回函数和行号
	mLog.SetReportCaller(global.Config.Logger.RowNumber)
	//自定义格式化
	mLog.SetFormatter(&LogFormatter{})
	//日志级别
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	return mLog
}
