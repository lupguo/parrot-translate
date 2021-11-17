package mlogger

import (
	"log"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

// Config 设置服务的日志存储文件
func Config(logFile string) middleware.LoggerConfig {
	openFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		log.Fatalf("set log file got err: %v", err)
	}

	// todo 日志滚动
	return middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           openFile,
	}
}
