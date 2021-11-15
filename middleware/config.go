package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BasicMiddlewares 服务配置的基础MiddleWare
var BasicMiddlewares = []echo.MiddlewareFunc{
	middleware.Logger(),
	middleware.Recover(),
}

// CustomMiddlewares 自定义Middleware
var CustomMiddlewares = []echo.MiddlewareFunc{
	AuthCheck,
	UserServiceLevel,
}
