package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := "./log"
	logFileName := "go_src.log"

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// 实例化
	logger := log.New()
	// 设置输出
	logger.Out = src
	logger.SetReportCaller(true)
	// 设置日志级别
	logger.SetLevel(log.DebugLevel)
	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		bodyJson := ""
		if c.Request.Body != nil {
			// https://zhuanlan.zhihu.com/p/87298102
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			bodyJson = string(bodyBytes)
		}
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logger.WithFields(log.Fields{
			"latency_time": latencyTime,
			"client_ip":    c.ClientIP(),
			"req_method":   c.Request.Method,
			"req_uri":      c.Request.RequestURI,
			"body_json":    bodyJson,
		}).Debug("记录请求")
	}
}
