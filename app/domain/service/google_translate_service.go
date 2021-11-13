package service

type ITranslate interface {
	// TranslateText 将text从sourceLang翻译成targetLang，返回值为翻译结果，有错直接返回
	TranslateText(sourceLang string, targetLang string, text string) (string, error)
}

type GoogleTranslateService struct {

}

func (g *GoogleTranslateService) TranslateText(sourceLang string, targetLang string, text string) (string, error) {
	panic("implement me")
}
