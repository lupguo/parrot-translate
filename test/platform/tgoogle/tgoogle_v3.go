package tgoogle

import (
	"context"
	"fmt"
	"io"

	translate "cloud.google.com/go/translate/apiv3"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// TranslateText translates input text and returns translated text.
func TranslateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string) error {
	// projectID := "my-project-id"
	// sourceLang := "en-US"
	// targetLang := "fr"
	// text := "Text you wish to translate"

	// explicit
	ctx := context.Background()
	jsonPath := `/data/projects/github.com/lupguo/parrot-translate/cmd/parrot-translate/sage-ace-331915-d22cf04c186b.json`

	// do translate
	// ctx := context.Background()
	clientx, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		return fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer clientx.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := clientx.TranslateText(ctx, req)
	if err != nil {
		return fmt.Errorf("TranslateTextV2: %v", err)
	}

	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
		fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
	}

	return nil
}
