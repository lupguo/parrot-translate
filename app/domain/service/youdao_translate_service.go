package service

import (
	"context"
)

type YoudaoTranslateService struct {

}

func (y *YoudaoTranslateService) TranslateText(ctx context.Context, srcLang string, toLang string, srcText string) (string, error) {
	panic("implement me")
}

