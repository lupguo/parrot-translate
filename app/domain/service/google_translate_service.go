package service

import (
	"context"

	"github.com/lupguo/parrot-translate/app/domain/repository"
)

type ITranslateService interface {
	// TranslateText 将text从sourceLang翻译成targetLang，返回值为翻译结果，有错直接返回
	TranslateText(ctx context.Context, srcLang string, toLang string, srcText string) (toText string, err error)
}

type GoogleTranslateService struct {
	repos repository.ITranslateRepos
}

func NewGoogleTranslateService(repos repository.ITranslateRepos) *GoogleTranslateService {
	return &GoogleTranslateService{repos: repos}
}

func (g *GoogleTranslateService) TranslateText(ctx context.Context, srcLang string, toLang string, srcText string) (toText string, err error) {
	return g.repos.GoogleTranslateText(ctx, srcLang, toLang, srcText)
}
