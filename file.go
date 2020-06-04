package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDir(fileName string) (err error) {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
	}
	return err
}

func CreateFile(path, content string) (result CreateFileResult) {

	result.path = path

	err := EnsureDir(path)

	if err == nil {
		f, err := os.Create(path)

		defer func() {
			if err2 := f.Close(); err2 != nil && err == nil {
				err = err2
				result.error = err
			}
		}()

		_, err = f.WriteString(content)
	}

	result.error = err

	return result

}

func CreateDir(dirPath string) (err error) {
	src, err := os.Stat(dirPath)

	if err == nil && src.Mode().IsRegular() {
		err = errors.New(fmt.Sprintf("%s already exist as a file.", dirPath))
		return err
	}

	if !os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("%s already exist as a dir.", dirPath))
	}

	err = os.MkdirAll(dirPath, 0755)

	return err
}
