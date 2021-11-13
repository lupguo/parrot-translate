package tgoogle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateTextV2(t *testing.T) {
	type args struct {
		targetLanguage string
		text           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t1", args{"zh-CN","hello world"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TranslateTextV2(tt.args.targetLanguage, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateTextV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got:%s", got)
			assert.NotNil(t, got)
		})
	}
}
