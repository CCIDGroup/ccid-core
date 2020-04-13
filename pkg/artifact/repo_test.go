package artifact

import (
	"fmt"
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

func Test_plainClone(t *testing.T) {
	type args struct {
		repo *Repo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"test plainClone",
			args {
				&Repo{
					"https://github.com/go-git/go-git.git",
					"master",
					"",
					"",
					"",
					"",
					"",
				},
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := plainClone(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("plainClone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(got)
		})
	}
}