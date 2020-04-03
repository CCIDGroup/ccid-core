/*
 * Copyright 2020 The CCID Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http: //www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package container

import (
	"bytes"
	"context"
	"github.com/CCIDGroup/ccid-core/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"runtime"
)

const (
	macPath     = "~/Library/Containers/com.docker.docker/Data/vms/0/"
	linuxPath   = "/var/lib/docker"
	windowsPath = "C:\\ProgramData\\DockerDesktop"
	unit        = "MB"
)

var e = &Engine{}

type Engine struct {
	Ctx      context.Context
	Instance *client.Client
}

func init() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	e.Instance = cli
	e.Ctx = ctx
}

//获取docker相关信息
func GetDockerEngineInfo() (*CheckList, error) {
	cl := &CheckList{}
	ver, err := e.Instance.ServerVersion(e.Ctx)
	if err != nil {

		return nil, err
	}
	cl.DockerEngineVersion = ver.Version
	cl.IsAvailable = true
	sysType := runtime.GOOS

	switch sysType {
	case "linux":
		cl.ImageStorePath = linuxPath
		cl.ContainerStorePath = linuxPath
	case "windows":
		cl.ImageStorePath = windowsPath
		cl.ContainerStorePath = windowsPath
	case "darwin":
		cl.ImageStorePath = macPath
		cl.ContainerStorePath = macPath
	default:
		cl.ContainerStorePath = "unknown"
		cl.ImageStorePath = "unknown"
	}
	cl.DiskSpaceUnit = unit
	cl.FreeDiskSpace = utils.GetFreeDiskSpace(cl.ImageStorePath)

	return cl, err
}

func PullImage(image string) error {

	reader, err := e.Instance.ImagePull(e.Ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	rev := utils.GenerateTaskID()
	l := (&utils.Log{}).InitLog()
	l.LogStream(rev, reader)
	return nil
}

func RunContainer(c *ContainerOpr) error {

	resp, err := e.Instance.ContainerCreate(e.Ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		return err
	}

	if err := e.Instance.ContainerStart(e.Ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := e.Instance.ContainerWait(e.Ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := e.Instance.ContainerLogs(e.Ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	//stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)

	return nil
}
