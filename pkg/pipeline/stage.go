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
package pipeline

import (
	"bytes"
	"fmt"
	"github.com/CCIDGroup/ccid-core/pkg/artifact"
	"github.com/CCIDGroup/ccid-core/pkg/docker"
	"github.com/CCIDGroup/ccid-core/pkg/message"
	"github.com/CCIDGroup/ccid-core/utils"
	"gopkg.in/jeevatkm/go-model.v1"
	"io/ioutil"
	"strings"
	"text/template"
)

type Stage struct {
	Name string                   `yaml:"name"`
	DisplayName string            `yaml:"displayName"`
	Container   Container         `yaml:"container"`
	Repository  Repository        `yaml:"repository"`
	Variables   map[string]string `yaml:"variables"`
	Jobs        []Job             `yaml:"jobs"`
}

func (s *Stage) RunStage(p *Run)(*chan string,error){

	repo,err := s.Repository.plainClone(p.Pipeline.PipelineID,p.RunID)
	s.Container.CodePath = strings.Replace(repo.FullPath,utils.GetCurrentDirectory()+p.Pipeline.PipelineID+"/"+p.RunID+"/","/tmp/ccid/",1)

	var m docker.Model
	model.Copy(&m, s.Container)
	m.HostCodePath = repo.FullPath
	if err != nil {
		return nil,err
	}
	var ch *chan string
	ch,err = docker.PullImage(&m)
	if err != nil {
		utils.LogError(err,"error to pull image")
		return nil, err
	}

	rmq := (&message.RabbitMQ{}).InitAMQP(p.RunID)

	rmq.WriteToMQ(ch)

	_,err = docker.CreateContainer(&m)
	if err != nil {
		return nil, err
	}

	err = docker.StartContainer(&m)

	if err != nil {
		return nil, err
	}

	for _,job := range s.Jobs {
		for _,step := range job.Steps {
			a := &Args{
				CodePath:     s.Container.CodePath,
				PipelineID:   p.Pipeline.PipelineID,
				RunID:        p.RunID,
				PipelineName: p.Pipeline.PipelineName,
				RunName:      p.RunName,
			}
			a.HostCodePath = repo.FullPath
			for _, scr := range step.Script{
				scr = hostCmd(scr,rmq,a)
				if scr == "" {
					continue
				}

				scr = parseScript(scr,a)
				ch,err = docker.ExecContainer(&m, scr)
				if err != nil {
					utils.LogError(err,"Exec container error")
				}
				rmq.WriteToMQ(ch)
			}
		}
	}
	rmq.Dispose()

	return ch,nil
}

func parseScript(script string, a *Args) string {
	t := template.New("scriptTemplate")
	t, _ = t.Parse(script)
	buf := new(bytes.Buffer)
	a.CurrentTime = utils.GenerateTaskID()
	t.Execute(buf,a)
	return buf.String()
}

func hostCmd(cmd string, rmq *message.RabbitMQ, a *Args) string {
	cmd = parseScript(cmd,a)
	cmds := strings.Split(cmd," ")
	switch cmds[0] {
	case "pkg":
		target,err := utils.ZipFolder(cmds[1],cmds[2])
		fmt.Println(target)
		if err != nil {
			fmt.Println(err.Error())
		}
		var so = &artifact.StorageOption{
			Endpoint:        "localhost:9000",
			AccessKeyID:     "minio",
			SecretAccessKey: "password01!",
			UseSSl:          false,
			BucketName:      utils.GenerateTaskID("20060102"),
			Location:        "us-east-1",
		}
		var s = &artifact.Storage{}
		s.InitStorage(so)
		s.UploadArtifact(cmds[2],target)
	case "image":
		filepath := cmds[1]
		filepath = strings.Replace(filepath,"//","/",-1)
		content,err := ioutil.ReadFile(filepath)
		df := parseScript(string(content),a)
		ioutil.WriteFile(filepath,[]byte(df),0777)
		ch,err := docker.BuildAndPushImage(filepath)
		if err != nil {
			utils.LogError(err,"Exec container error")
		}
		rmq.WriteToMQ(ch)
	default:
		return cmd
	}
	return ""
}







