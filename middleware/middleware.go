package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lupguo/parrot-translate/middleware/authcheck"
	"github.com/lupguo/parrot-translate/middleware/mlogger"
	"github.com/spf13/viper"
)

// InitConfig 初始化中间配置
func InitConfig(e *echo.Echo, v *viper.Viper) {
	// logger config
	logFile := v.GetString("plugin.log.log_file")
	logger := middleware.LoggerWithConfig(mlogger.Config(logFile))

	// setting middleware
	e.Use(
		middleware.Recover(),
		logger,
	)
}

// BasicMiddlewares 服务配置的基础MiddleWare
var BasicMiddlewares = []echo.MiddlewareFunc{}

// CustomMiddlewares 自定义Middleware
var CustomMiddlewares = []echo.MiddlewareFunc{
	authcheck.AuthCheck,
	authcheck.UserServiceLevel,
}
