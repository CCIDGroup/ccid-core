package artifact

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/CCIDGroup/ccid-core/utils"
	"github.com/go-git/go-git/v5"
	"github.com/rs/xid"
	"strings"
)

const (
	CodePath = "app/"
)

type Repo struct {
	URL string
	UserName string
	Password string
	Folder   string
	FullPath string
	Message  string
}

func plainClone(repo *Repo) (*Repo, error) {
	path := utils.GetCurrentDirectory()
	fullPath := path + CodePath
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

	if !strings.HasPrefix(repo.URL, "https://") {
		return repo, errors.New("wrong git prefix, should be start with https")
	}
	url := repo.URL
	if repo.UserName != "" && repo.Password != "" {
		relativeUrl := strings.ReplaceAll(repo.URL,"https://","")
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
