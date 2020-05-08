package docker

import (
	"fmt"
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
		c *Model
	}
	tests := []struct {
		name    string
		args    args
		want    *chan string
		wantErr bool
	}{
		{
			"dotnetsdk",
			args{
				&Model{
					ID:       "",
					Image:    "sdk:3.1-alpine3.11",
					Endpoint: "mcr.microsoft.com/dotnet/core",
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
			got, err := PullImage(tt.args.c)
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
		c *Model
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
				c: &Model{
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
			id, err := CreateContainer(tt.args.c,"/")
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
		c *Model
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestStartContainer",
			args{
				&Model{
					ID: ID,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartContainer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("StartContainer() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
	time.Sleep(time.Second*5)
}

func TestExecContainer(t *testing.T) {
	type args struct {
		c *Model
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
				&Model{
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
			got, err := ExecContainer(tt.args.c,tt.args.scripts)
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
		c *Model
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
				&Model{
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