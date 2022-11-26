package config

import (
	"github.com/mailpiggy/MailPiggy/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Smtp struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"smtp"`
	Http struct {
		Host string `yaml:"host"`
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
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
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
	if len(config.Smtp.Host) <= 0 {
		config.Smtp.Host = "0.0.0.0"
	}
	if config.Smtp.Port <= 0 {
		config.Smtp.Port = 1025
	}

	if len(config.Http.Host) <= 0 {
		config.Http.Host = "0.0.0.0"
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

	if len(config.Log.Level) <= 0 {
		config.Log.Level = logger.DEBUG
	}
}
