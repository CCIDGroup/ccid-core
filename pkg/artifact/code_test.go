package artifact

import (
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test clone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}


}