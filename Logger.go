package Logger

import (
	log "github.com/sirupsen/logrus"
	"path"
	"runtime"
	"fmt"
)

func InitLogger() (Logger *log.Logger) {
	Logger = log.New()
	Logger.SetLevel(log.DebugLevel)
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		ForceQuote:      true,
		TimestampFormat: "2006/01/02 15:04:05:06",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// 剪切调用logrus的函数路径
			fn := path.Base(frame.File)
			// 拼接路径，行数
			fileName := fmt.Sprintf(" %s:%d",fn,frame.Line)
			// 拼接调用函数
			fc := fmt.Sprintf("%s()",frame.Function)
			return fc, fileName
		},
	})
	return
}



