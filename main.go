package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/parrot-translate/app/application"
	"github.com/lupguo/parrot-translate/app/domain/service"
	"github.com/lupguo/parrot-translate/app/handler"
	"github.com/lupguo/parrot-translate/app/infrastruct/cache"
	"github.com/lupguo/parrot-translate/app/infrastruct/translate"
	"github.com/lupguo/parrot-translate/cmd"
	"github.com/lupguo/parrot-translate/middleware"
)

func main() {
	// 命令行解析
	if err := cmd.Execute(); err != nil {
		log.Fatalf("command line execute got err:%v", err)
	}
	cfg := cmd.GetConfigViper()

	// 依赖基础设施初始化
	projectID := cfg.GetString("api.google.proj_id")
	jsonPath := cfg.GetString("api.google.auth_file")
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
	e.Use(middleware.BasicMiddlewares...)

	// 服务启动
	e.Logger.Fatal(e.Start(cfg.GetString("server.listen")))
}
