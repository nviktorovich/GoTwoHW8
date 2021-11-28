package main

import (
	"fmt"
	"github.com/nviktorovich/copyhunter/program"
	"github.com/nviktorovich/copyhunter/program/config"
	"os"
)

func main() {
	cfg := config.Config{}

	if err := cfg.GetConfig(); err != nil {
		fmt.Println("ошибка чтения из конфигурации, ", err)
		fmt.Println("загружена конфигурация по умолчанию")
		cfg.GetDefaultConfig()
	}
	if cfg.Switcher == 0 {
		cfg.GetDefaultConfig()
	}
	if err := cfg.ConfigPrinter(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileMap, err := program.CreateFilesNamesList(cfg.Root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	program.FileRemover(fileMap, cfg)

}
