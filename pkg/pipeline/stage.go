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
	"github.com/CCIDGroup/ccid-core/pkg/docker"
	"github.com/CCIDGroup/ccid-core/pkg/message"
	"github.com/CCIDGroup/ccid-core/utils"
	"gopkg.in/jeevatkm/go-model.v1"
	"strings"
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
	repo,err := s.Repository.plainClone(p.pipeline.pipelineID,p.runID)
	s.Container.CodePath = strings.Replace(repo.FullPath,utils.GetCurrentDirectory(),"/tmp/ccid/",0)

	var m docker.Model
	model.Copy(&m, s.Container)
	if err != nil {
		return nil,err
	}
	var ch *chan string
	ch,err = docker.PullImage(&m)
	if err != nil {
		utils.LogError(err,"error to pull image")
		return nil, err
	}

	rmq := (&message.RabbitMQ{}).InitAMQP(p.runID)

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
			ch,err = docker.ExecContainer(&m, step.Script)
			if err != nil {
				utils.LogError(err,"Exec container error")
			}
			rmq.WriteToMQ(ch)
		}
	}

	return ch,nil
}





