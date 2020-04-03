package container

import (
	"testing"
)

func TestGetDockerEngineInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    *CheckList
		wantErr bool
	}{
		{
			name:    "Test Docker Info",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDockerEngineInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDockerEngineInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				t.Log(got)
			}
		})
	}
}

func TestPullImage(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"minio/minio",
			args{"minio/minio"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PullImage(tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("PullImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
