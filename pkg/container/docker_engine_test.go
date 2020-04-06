package container

import (
	"fmt"
	"github.com/CCIDGroup/ccid-core/utils"
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
		rev   string
		image string
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
				utils.GenerateTaskID(),
				"minio/minio",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PullImage(tt.args.image)
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
		c *ConOpr
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
				c: &ConOpr{
					Name:     "c_201012121212",
					Image:    "minio/minio",
					Endpoint: "",
					Env:      []string{},
					Cmd:      []string{"server", "/data"},
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
			id, err := CreateContainer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(id)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("StartContainer() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestStartContainer(t *testing.T) {
	type args struct {
		c *ConOpr
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestStartContainer",
			args{
				&ConOpr{
					ID: "1890d33e847630f4f4d0b88a0e2746fe54055b219d7530439c6205771cb296ce",
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
}

func TestLogContainer(t *testing.T) {
	type args struct {
		c *ConOpr
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
				&ConOpr{
					ID: "1890d33e847630f4f4d0b88a0e2746fe54055b219d7530439c6205771cb296ce",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogContainer(tt.args.c)
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

func TestExecContainer(t *testing.T) {
	type args struct {
		c *ConOpr
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
				&ConOpr{
					ID: "1890d33e847630f4f4d0b88a0e2746fe54055b219d7530439c6205771cb296ce",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecContainer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecContainer() error = %v, wantErr %v", err, tt.wantErr)
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
