package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline_Create(t *testing.T) {
	type fields struct {
		Name         string
		Variables    map[string]string
		Trigger      Trigger
		Stages       []Stage
		pipelineName string
		pipelineID   string
		runName      string
		runID        string
		sourceBranch string
		sourceCommit string
	}
	type args struct {
		pipeline string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pipeline
		wantErr bool
	}{
		{
			"Start Pipeline",
			fields{},
			args{
				"name: TEST\nstages:\n- name: \"stage1\"\n  displayName: \"test stage\"\n  container:\n    image: \"sdk:3.1-alpine3.11\"\n    endpoint: \"mcr.microsoft.com/dotnet/core\"\n  repository:\n    type: Github\n    name: \"dotnet demo\"\n    ref: master \n    endpoint: \"https://github.com/CCIDGroup/ccid-dotnet-sample.git\"\n  jobs:\n  - name: build\n    displayName: \"build dotnet\"\n    steps:\n    - script:\n      - \"dotnet restore\"\n      - \"dotnet build --output {{.CodePath}}pkg/\"\n      - \"pkg {{.HostCodePath}}/pkg/ {{.CurrentTime}}.zip\"\n      - \"image {{.HostCodePath}}/Dockerfile\"",
			},
			nil,
			false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pipeline{
				Name:         tt.fields.Name,
				Variables:    tt.fields.Variables,
				Trigger:      tt.fields.Trigger,
				Stages:       tt.fields.Stages,
				PipelineName: tt.fields.pipelineName,
				PipelineID:   tt.fields.pipelineID,
				sourceBranch: tt.fields.sourceBranch,
				sourceCommit: tt.fields.sourceCommit,
			}
			_, err := p.Create(tt.args.pipeline)
			if err != nil {
				fmt.Println(err.Error())
			}

		})
	}
}