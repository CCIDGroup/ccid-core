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
	"fmt"
	"github.com/rs/xid"
	"gopkg.in/yaml.v2"
	"time"
)

type Pipeline struct {
	Name         string            `yaml:"name"`
	Variables    map[string]string `yaml:"variables"`
	Trigger      Trigger           `yaml:"trigger"`
	Stages       []Stage           `yaml:"stages"`
	PipelineName string
	PipelineID   string
	sourceBranch string
	sourceCommit string
}

type Run struct {
	RunID       string
	RunName     string
	Pipeline    *Pipeline
}

func (p *Pipeline) Create(pipeline string) (*Run,error){
	err := yaml.Unmarshal([]byte(pipeline), p)
	if err != nil {
		return nil, err
	}
	p.PipelineID = xid.New().String()
	r := &Run{}
	r.RunID = xid.New().String()
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	r.RunName = fmt.Sprintf("pipeline-%vrun-%v start time:%v",p.PipelineID,r.RunID,dateTime)
	r.Pipeline = p
	for _,stage := range p.Stages {
		stage.RunStage(r)
	}
	return r,nil
}







