package translate

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// GoogleTranslateInfra Google翻译基础设施
type GoogleTranslateInfra struct {
	projectID string
	jsonPath  string
}

func NewGoogleTranslateInfra(projectID string, jsonPath string) *GoogleTranslateInfra {
	return &GoogleTranslateInfra{
		projectID: projectID,
		jsonPath:  jsonPath,
	}
}

func (g *GoogleTranslateInfra) GoogleTranslateText(ctx context.Context, sourceLang string, targetLang string, sourceText string) (toText string, err error) {
	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(g.jsonPath))
	if err != nil {
		return "", errors.Wrap(err, "GoogleTranslateInfra > translate.NewTranslationClient() got err")
	}
	defer client.Close()

	// 请求格式
	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", g.projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{sourceText},
	}

	// 发送请求
	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, "GoogleTranslateInfra > client.TranslateText() got err")
	}

	// Display the translation for each input text provided
	var texts []string
	for _, translation := range resp.GetTranslations() {
		texts = append(texts, translation.GetTranslatedText())
	}

	return texts[0], nil
}
