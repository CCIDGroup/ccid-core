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
	"errors"
	"fmt"
	"github.com/CCIDGroup/ccid-core/utils"
	"github.com/go-git/go-git/v5"
	"github.com/rs/xid"
	"strings"
)

type RepositoryType string

const (
	Github RepositoryType = "Github"
	Gitea  RepositoryType = "Gitea"
	CodePath = "/app/"

)

type Repository struct {
	RepositoryType RepositoryType `yaml:"type"`
	Name string `yaml:"name"`
	Ref  string `yaml:"ref"`
	Endpoint string `yaml:"endpoint"`
	URL string
	Branch   string
	UserName string
	Password string
	Folder   string
	FullPath string
	Message  string
}


func (repo *Repository) plainClone(pID,runID string) (*Repository, error) {
	path := utils.GetCurrentDirectory()
	fullPath := path + pID+"/"+runID + CodePath
	if !utils.Exist(fullPath){
		utils.CreateDir(fullPath)
	}
	xid := xid.New().String()
	filePath := fullPath + xid
	if !utils.Exist(filePath){
		utils.CreateDir(filePath)
	}

	repo.Folder = xid
	repo.FullPath = filePath

	if !strings.HasPrefix(repo.Endpoint, "https://") {
		return repo, errors.New("wrong git prefix, should be start with https")
	}
	url := repo.Endpoint
	if repo.UserName != "" && repo.Password != "" {
		relativeUrl := strings.ReplaceAll(repo.Endpoint,"https://","")
		url = fmt.Sprintf("https://%s:%s@%s", repo.UserName, repo.Password, relativeUrl)
	}
	buf := new(bytes.Buffer)
	_, err := git.PlainClone(filePath, false, &git.CloneOptions{
		URL:      url,
		Progress: buf,
	})
	if err != nil {
		return repo, err
	}
	repo.Message = buf.String()
	return repo, nil
}
