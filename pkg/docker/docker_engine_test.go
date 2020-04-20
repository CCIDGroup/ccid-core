package docker

import (
	"fmt"
	"github.com/CCIDGroup/ccid-core/pkg/pipeline"
	"testing"
	"time"
)
var ID string

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
			got, err := getDockerEngineInfo()
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
		c *pipeline.Container
	}
	tests := []struct {
		name    string
		args    args
		want    *chan string
		wantErr bool
	}{
		{
			"minio/minio",
			args{
				&pipeline.Container{
					ID:       "",
					Name:     "minio/minio",
					Image:    "minio/minio",
					Endpoint: "",
					Env:      nil,
					Cmd:      nil,
					Options:  "",
					Ports:    nil,
					Volumes:  nil,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pullImage(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for {
				val, ok := <-*got
				if ok == false {
					fmt.Println("pull image done")
					break
				} else {
					fmt.Print(val)
				}
			}
		})
	}
}

func TestCreateContainer(t *testing.T) {

	type args struct {
		c *pipeline.Container
	}
	tests := []struct {
		name    string
		args    args
		want    *chan string
		wantErr bool
	}{
		{
			"start minio",
			args{
				c: &pipeline.Container{
					Name:     "c_201012121212",
					Image:    "minio/minio",
					Endpoint: "",
					Env:      []string{},
					Cmd:      []string{"server","/data"},
					Options:  "",
					Ports:    []string{"9000:9000"},
					Volumes:  []string{},
				},
			},
			nil,
			false,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := createContainer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ID  = id
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("StartContainer() got = %v, want %v", got, tt.want)
			//}
		})
	}


}

func TestStartContainer(t *testing.T) {
	type args struct {
		c *pipeline.Container
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestStartContainer",
			args{
				&pipeline.Container{
					ID: ID,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startContainer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("StartContainer() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
	time.Sleep(time.Second*5)
}

func TestExecContainer(t *testing.T) {
	type args struct {
		c *pipeline.Container
		scripts []string
	}
	tests := []struct {
		name    string
		args    args
		want    *chan string
		wantErr bool
	}{
		{
			"TestLogContainer",
			args{
				&pipeline.Container{
					ID: ID,
				},
				[]string{"echo","hello world"},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execContainer(tt.args.c,tt.args.scripts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for {
				val, ok := <-*got
				if ok == false {
					fmt.Println("exec docker done")
					break
				} else {
					fmt.Print(val)
				}
			}

		})
	}
}

func TestLogContainer(t *testing.T) {
	type args struct {
		c *pipeline.Container
	}
	tests := []struct {
		name    string
		args    args
		want    *chan string
		wantErr bool
	}{
		{
			"TestLogContainer",
			args{
				&pipeline.Container{
					ID: ID,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := logContainer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for {
				val, ok := <-*got
				if ok == false {
					fmt.Println("log done")
					break
				} else {
					fmt.Print(val)
				}
			}
		})
	}

}