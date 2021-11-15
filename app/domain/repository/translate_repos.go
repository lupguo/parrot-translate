package repository

import (
	"context"
)

type ITranslateRepos interface {
	GoogleTranslateText(ctx context.Context, sourceLang string, targetLang string, sourceText string) (toText string, err error)
}
