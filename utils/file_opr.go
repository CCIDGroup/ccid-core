package utils

import (
	"io"
	"os"
)

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(file string) (io.Writer, error) {
	if Exist(file) {
		return os.Open(file)
	} else {
		return os.Create(file)
	}

}
