package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lupguo/parrot-translate/app/application"
)

type TranslateHandler struct {
	app *application.TranslateApp
}

// NewTranslateHandler 初始化业务处理程序
func NewTranslateHandler(app *application.TranslateApp) *TranslateHandler {
	return &TranslateHandler{app: app}
}

// TranslateText 翻译接口
func (h *TranslateHandler) TranslateText(c echo.Context) error {
	// 请求ctx
	ctx := context.Background()
	context.WithValue(ctx, "c", c)
	context.WithDeadline(ctx, time.Now().Add(3*time.Second))

	// 请求参数
	srcLang := c.FormValue("srcLang")
	toLang := c.FormValue("toLang")
	text := c.FormValue("text")

	// 业务接口处理
	ent, err := h.app.TranslateText(ctx, srcLang, toLang, text)
	if err != nil {
		c.Logger().Errorf("h.app.TranslateText() got err: %s", err)
		return c.JSON(http.StatusInternalServerError, "translate text got error, try again latter!")
	}

	// 数据转换dto
	return c.JSON(http.StatusOK, ent)
}
