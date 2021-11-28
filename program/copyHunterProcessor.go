package program

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nviktorovich/copyhunter/program/config"
	"os"
)

func Remover(f FileStruct) error {
	if err := os.Remove(f.FileLocation); err != nil {
		err = errors.New(fmt.Sprintf("ошибка удаления файла %v", f.FileLocation))
		return err
	}
	return nil
}

func FileRemover(filesMap []FileStruct, cfg config.Config) error {
	origin := make(map[string]FileStruct)
	fmt.Println(cfg)
	for _, v := range filesMap {
		_, ok := origin[v.FileLocation]
		if !ok {
			for _, holdFile := range origin {
				if v.Info.Name() == holdFile.Info.Name() && cfg.Mode == 0 {
					if err := Remover(v); err != nil {
						return err
					}
				}
				if bytes.Compare(v.Content, holdFile.Content) == 0 && cfg.Mode == 1 {
					if err := Remover(v); err != nil {
						return err
					}
				}
				if (v.Info.Name() == holdFile.Info.Name()) && (bytes.Compare(v.Content, holdFile.Content) == 0) && cfg.Mode == 2 {
					if err := Remover(v); err != nil {
						return err
					}
				}
			}
			origin[v.FileLocation] = v
		}
	}
	fmt.Println(origin)
	return nil
}
