package config

import (
	"crypto/rand"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Hostname string `yaml:"hostname"`
	Smtp     struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Validation struct {
			MaximumLineLength int `yaml:"maximum_line_length"`
			MaximumReceivers  int `yaml:"maximum_receivers"`
		} `yaml:"validation"`
	} `yaml:"smtp"`
	Http struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Path         string `yaml:"path"`
		AllowOrigins string `yaml:"allow_origins"`
		AssetsRoot   string `yaml:"assets_root"`
		BaseUrl      string `yaml:"base_url"`
	} `yaml:"http"`
	Websocket struct {
		MaxConnection int `yaml:"max_connection"`
	} `yaml:"websocket"`
	Storage struct {
		Use          string `yaml:"use"`
		PerRoomLimit int    `yaml:"per_room_limit"`
		Directory    struct {
			Path string `yaml:"path"`
		} `yaml:"directory"`
		MongoDB struct {
			Connection string `yaml:"connection"`
			Collection string `yaml:"collection"`
		} `yaml:"mongodb"`
	} `yaml:"storage"`
	Authentication struct {
		Use      string `yaml:"use"`
		Type     string `yaml:"type"`
		Admin    string `yaml:"admin"`
		Internal struct {
			HmacSecret []byte
			// Lifetime in minutes
			SessionLifetime int64 `yaml:"session_lifetime"`
		} `yaml:"internal"`

		File struct {
			Path string `yaml:"path"`
		} `yaml:"file"`
		MongoDB struct {
			Connection string `yaml:"connection"`
			Collection string `yaml:"collection"`
		} `yaml:"mongodb"`
	} `yaml:"authentication"`
	UI struct {
		File string `yaml:"file"`
	} `yaml:"ui"`
	DB  DbConfig `yaml:"db"`
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
	if len(config.Hostname) <= 0 {
		identification, _ := os.Hostname()
		if len(identification) <= 0 {
			config.Hostname = "mailhedgehog.local"
		} else {
			config.Hostname = identification
		}
	}

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

	if config.Websocket.MaxConnection <= 0 {
		config.Websocket.MaxConnection = 50
	}

	if len(config.Storage.Use) <= 0 {
		config.Storage.Use = "directory"
	}

	if len(config.Storage.Directory.Path) <= 0 {
		config.Storage.Directory.Path = ""
	}

	if len(config.Storage.MongoDB.Connection) <= 0 {
		config.Storage.MongoDB.Connection = "mongodb"
	}

	if len(config.Storage.MongoDB.Collection) <= 0 {
		config.Storage.MongoDB.Collection = "emails"
	}

	if len(config.Authentication.Use) <= 0 {
		config.Authentication.Use = "file"
	}

	if len(config.Authentication.Type) <= 0 {
		config.Authentication.Type = "internal"
	}

	if len(config.Authentication.File.Path) <= 0 {
		config.Authentication.File.Path = ""
	}

	if len(config.Authentication.MongoDB.Connection) <= 0 {
		config.Authentication.MongoDB.Connection = "mongodb"
	}

	if len(config.Authentication.MongoDB.Collection) <= 0 {
		config.Authentication.MongoDB.Collection = "users"
	}

	if len(config.Log.Level) <= 0 {
		config.Log.Level = logger.DEBUG
	}

	if config.Authentication.Type == "internal" {
		hmacSecret := make([]byte, 200)
		_, err := rand.Read(hmacSecret)
		logger.PanicIfError(err)
		config.Authentication.Internal.HmacSecret = hmacSecret

		if config.Authentication.Internal.SessionLifetime <= 0 {
			config.Authentication.Internal.SessionLifetime = 2 * 60
		}
	}
}

func (config *AppConfig) IsAdmin(username string) bool {
	return len(config.Authentication.Admin) > 0 && config.Authentication.Admin == username
}
