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
package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"runtime"
	"strings"
	"time"
)

const (
	macPath      = "~/Library/Containers/com.docker.docker/Data/vms/0/"
	linuxPath    = "/var/lib/docker"
	windowsPath  = "C:\\ProgramData\\DockerDesktop"
	relativePath = "/tmp/ccid/"
	unit         = "MB"
)

var ctx      context.Context
var instance *client.Client

func init() {
	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	instance = cli

}

//获取docker相关信息
func GetDockerEngineInfo() (*CheckList, error) {
	cl := &CheckList{}
	ver, err := instance.ServerVersion(ctx)
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
	//cl.FreeDiskSpace = utils.GetFreeDiskSpace(cl.ImageStorePath)

	return cl, err
}

func PullImage(c *Model) (*chan string, error) {
	image := c.Image
	if c.Endpoint != "" {
		image = c.Endpoint + "/" + image
	}

	reader, err := instance.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	ch := logStream(reader)
	return ch, nil
}

func CreateContainer(c *Model) (string, error) {
	image := c.Image
	if c.Endpoint != "" {
		image = c.Endpoint + "/" + image
	}

	exposedPorts, portBindings, _ := nat.ParsePortSpecs(c.Ports)
	c.Volumes = append(c.Volumes, c.CodePath)
	resp, err := instance.ContainerCreate(ctx, &container.Config{
		Image:        image,
		Env:          c.Env,
		ExposedPorts: exposedPorts,
		Cmd:          c.Cmd,
		Tty:          true,
	}, &container.HostConfig{
		Binds:        c.Volumes,
		PortBindings: portBindings,
	}, nil, "")
	c.ID = resp.ID
	return c.ID, err
}

func StartContainer(c *Model) error {
	if err := instance.ContainerStart(ctx, c.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func StopContainer(c *Model) error {
	timeout := time.Second * 1
	if err := instance.ContainerStop(ctx, c.ID, &timeout); err != nil {
		return err
	}
	return nil
}

func RemoveContainer(c *Model) error {
	if err := instance.ContainerRemove(ctx, c.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	}); err != nil {
		return err
	}
	return nil
}

func ExecContainer(c *Model, script string) (*chan string, error) {
	exec, err := instance.ContainerExecCreate(ctx, c.ID, types.ExecConfig{
		User:         "",
		Privileged:   true,
		Tty:          true,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Detach:       false,
		Env:          []string{},
		Cmd:          strings.Split(script," "),
	})
	if err != nil {
		return nil, err
	}
	execAttachConfig := types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	}
	containerConn, e := instance.ContainerExecAttach(ctx, exec.ID, execAttachConfig)
	if err != nil {
		return nil, e
	}
	return logStream(containerConn.Reader), nil
}

//func BuildAndPushImage(dockerfile,tag string)(*chan string, error) {
//	opt := types.ImageBuildOptions{
//		Dockerfile:   "image/centos7/Dockerfile",
//	}
//	resp, err := instance.ImageBuild(context.Background(), nil, opt)
//	if err == nil {
//		fmt.Printf("Error, %v", err)
//	}
//	options := types.ImagePushOptions{}
//
//	return logStream(resp.Body), nil
//
//
//}

func logContainer(c *Model) (*chan string, error) {
	reader, err := instance.ContainerLogs(ctx, c.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Since:      "",
		Until:      "",
		Timestamps: false,
		Follow:     false,
		Tail:       "",
		Details:    true,
	})
	if err != nil {
		return nil, err
	}
	return logStream(reader), nil
}
func logStream(reader io.Reader) *chan string {
	r := make(chan string)
	go func() {
		for {
			buf := make([]byte, 1024)
			// 循环读取文件
			n, err2 := reader.Read(buf)
			if err2 != nil {
				break
			}
			r <- string(buf[:n])
		}
		close(r)
	}()
	return &r
}
