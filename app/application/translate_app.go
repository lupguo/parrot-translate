package application

import (
	"context"

	"github.com/lupguo/parrot-translate/app/domain/entity"
	"github.com/lupguo/parrot-translate/app/domain/repository"
	"github.com/lupguo/parrot-translate/app/domain/service"
	"github.com/lupguo/parrot-translate/pkg/trans"
)

type TranslateApp struct {
	cache     repository.ICacheRepos
	transSrvs []service.ITranslateService
}

// NewTranslateApp 初始化翻译应用程序，通过注入rds仓储实现、服务实现完成
func NewTranslateApp(cache repository.ICacheRepos, transSrvs ...service.ITranslateService) *TranslateApp {
	return &TranslateApp{cache: cache, transSrvs: transSrvs}
}

// TranslateText 应用层检测原始sourceText的hashID，对比之前是否已有过翻译记录，并返回翻译结果
func (app *TranslateApp) TranslateText(ctx context.Context, srcLang string, toLang string, srcText string) (ent *entity.TransEntity, err error) {
	// 1. 检测sourceText的hashID值是否已存在
	sourceHashID, err := trans.HashText(srcText)
	if err != nil {
		return nil, err
	}
	_ = sourceHashID

	// 2. todo redis get and check

	// 3. not exist, do srv.translate
	transSrv := app.transSrvs[0]
	text, err := transSrv.TranslateText(ctx, srcLang, toLang, srcText)

	// 4. todo set to redis

	// 返回
	return &entity.TransEntity{
		HashID:     sourceHashID,
		Text:       text,
		Additional: "",
	}, err
}
