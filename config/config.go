package config

import (
	"github.com/mailpiggy/MailPiggy/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Smtp struct {
		Port int `yaml:"port"`
	} `yaml:"smtp"`
	Http struct {
		Port int    `yaml:"port"`
		Path string `yaml:"path"`
	} `yaml:"http"`
	Storage struct {
		Use          string `yaml:"use"`
		PerRoomLimit int    `yaml:"per_room_limit"`
		Directory    struct {
			Path string `yaml:"path"`
		} `yaml:"directory"`
	} `yaml:"storage"`
	Authorisation struct {
		Use  string `yaml:"use"`
		File struct {
			Path string `yaml:"path"`
		} `yaml:"file"`
	} `yaml:"authorisation"`
}

func ParseConfig(filePath string) *AppConfig {
	config := &AppConfig{}

	if len(filePath) > 0 {
		yamlFile, err := os.ReadFile(filePath)
		logger.PanicIfError(err)
		err = yaml.Unmarshal(yamlFile, config)
		logger.PanicIfError(err)
	}

	config.withDefaults()

	return config
}

func (config *AppConfig) withDefaults() {
	if config.Smtp.Port <= 0 {
		config.Smtp.Port = 1025
	}

	if config.Http.Port <= 0 {
		config.Http.Port = 8025
	}

	if len(config.Storage.Use) <= 0 {
		config.Storage.Use = "directory"
		if len(config.Storage.Directory.Path) <= 0 {
			config.Storage.Directory.Path = ""
		}
	}

	if len(config.Authorisation.Use) <= 0 {
		config.Authorisation.Use = "file"
		if len(config.Authorisation.File.Path) <= 0 {
			config.Authorisation.File.Path = ""
		}
	}
}
