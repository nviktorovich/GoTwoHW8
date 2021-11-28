package program

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileStruct struct {
	FileLocation string
	Info         os.FileInfo
	Content      []byte
}

func CreateFilesNamesList(root string) ([]FileStruct, error) {
	var files []FileStruct
	var data []byte
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data, err = os.ReadFile(path)
			if err != nil {
				err = errors.New(fmt.Sprintf("ошибка чтения файла, %v", path))
			}
			files = append(files, FileStruct{path, info, data})
		}
		return nil
	}); err != nil {
		err = errors.New(fmt.Sprintf("ошибка построения карты, %v", err))
		return files, err
	}
	return files, nil
}
