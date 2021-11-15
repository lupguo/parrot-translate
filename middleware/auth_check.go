package middleware

import (
	"github.com/labstack/echo/v4"
)

// AuthCheck 验签中间件
func AuthCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// auth 验签
		return next(c)
	}
}

func UserServiceLevel(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// auth 客户等级校验
		return next(c)
	}
}
