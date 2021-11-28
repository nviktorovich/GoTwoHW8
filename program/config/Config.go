package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Switcher int    `yaml:"switcher"`
	Mode     int    `yaml:"mode"`
	Log      int    `yaml:"log"`
	Root     string `yaml:"root"`
}

type CfgManager interface {
	GetConfig() error
	GetDefaultConfig()
	ConfigPrinter()
}

func (cfg *Config) GetConfig() (err error) {
	var configData []byte
	configData, err = os.ReadFile("program/config/configuration.yaml")

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(configData, &cfg); err != nil {
		err = errors.New("ошибка разбора конфигурационного файла")
		return err
	}

	return nil
}

func (cfg *Config) GetDefaultConfig() {
	cfg.Switcher = 0
	cfg.Mode = 0
	cfg.Log = 0
	cfg.Root = "testDirectory/"
}

func (cfg *Config) ConfigPrinter() error {
	switch cfg.Switcher {
	case 0:
		fmt.Println("switcher: ", cfg.Switcher, "режим по умолчанию, значения конфигурационного файла игнорируются")
	case 1:
		fmt.Println("switcher: ", cfg.Switcher, "чтение из конфигурационного файла")
	default:
		err := errors.New(fmt.Sprintf("параметр определен неверно, допустимые значения 0/1, установлено: %v", cfg.Switcher))
		return err
	}

	switch cfg.Mode {
	case 0:
		fmt.Println("mode: ", cfg.Mode, "чистка файлов по имени")
	case 1:
		fmt.Println("mode: ", cfg.Mode, "чистка по содержимому")
	case 2:
		fmt.Println("mode: ", cfg.Mode, "чистка по полному совпадению (имя + содержимое)")
	default:
		err := errors.New(fmt.Sprintf("параметр определен неверно, допустимые значения 0/1/2, установлено: %v", cfg.Mode))
		return err
	}

	switch cfg.Log {
	case 0:
		fmt.Println("log: ", cfg.Log, "лог-файл выключен")
	case 1:
		fmt.Println("log: ", cfg.Log, "лог-файл включен")
	default:
		err := errors.New(fmt.Sprintf("параметр определен неверно, допустимые значения 0/1, установлено: %v", cfg.Log))
		return err
	}
	fmt.Println("корневая папка: ", cfg.Root)
	return nil
}
