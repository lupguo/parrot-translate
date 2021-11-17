package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lupguo/parrot-translate/app/application"
	"github.com/lupguo/parrot-translate/app/domain/service"
	"github.com/lupguo/parrot-translate/app/handler"
	"github.com/lupguo/parrot-translate/app/infrastruct/cache"
	"github.com/lupguo/parrot-translate/app/infrastruct/translate"
	"github.com/lupguo/parrot-translate/middleware"
	"github.com/spf13/viper"
)

func Start(vc *viper.Viper) {
	// 依赖基础设施初始化
	projectID := vc.GetString("api.google.proj_id")
	jsonPath := vc.GetString("api.google.auth_file")
	redisInfra := cache.NewRedisInfra()
	googleHTTPTranslateInfra := translate.NewGoogleTranslateInfra(projectID, jsonPath)

	// 注入并完成服务组装
	h := handler.NewTranslateHandler(
		application.NewTranslateApp(
			redisInfra,
			service.NewGoogleTranslateService(googleHTTPTranslateInfra),
		),
	)

	// 服务启动
	e := echo.New()

	// 路由配置
	e.Any("/translate-text", h.TranslateText, middleware.CustomMiddlewares...)
	e.Static("/static", "static")

	// 通用中间件配置
	middleware.InitConfig(e, vc)

	// 服务启动
	e.Logger.Fatal(e.Start(vc.GetString("server.listen")))
}
