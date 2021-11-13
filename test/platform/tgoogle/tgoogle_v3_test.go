package tgoogle

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s1 = `A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is nil.`
)

func TestTranslateText(t *testing.T) {
	type args struct {
		projectID  string
		sourceLang string
		targetLang string
		text       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t1", args{"sage-ace-331915", "en", "zh-CN", "hello world"}, false},
		{"t2", args{"sage-ace-331915", "en", "zh-CN", s1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := TranslateText(w, tt.args.projectID, tt.args.sourceLang, tt.args.targetLang, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, w.String())
			t.Logf("got %s", w.String())
		})
	}
}
