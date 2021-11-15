package trans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashText(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t1", args{"hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHashID, err := HashText(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got hash id: %s", gotHashID)
			assert.NotNil(t, gotHashID)
		})
	}
}
