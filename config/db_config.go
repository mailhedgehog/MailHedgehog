package config

import (
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/db"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbConfig struct {
	Connections map[string]map[string]interface{} `yaml:"connections"`
}

func (dbConfig *DbConfig) GetMongoDBConnection(connectionName string) *mongo.Database {
	config := dbConfig.Connections[connectionName]
	if config == nil {
		logger.PanicIfError(errors.New(fmt.Sprintf("Undefined db connection [%s]", connectionName)))
	}

	uri := "127.0.0.1:27017"
	if config["uri"] != nil && len(config["uri"].(string)) > 0 {
		uri = config["uri"].(string)
	}

	dbName := "mailhedgehog"
	if config["db_name"] != nil && len(config["db_name"].(string)) > 0 {
		dbName = config["db_name"].(string)
	}

	dbUser := ""
	if config["db_user"] != nil && len(config["db_user"].(string)) > 0 {
		dbUser = config["db_user"].(string)
	}

	dbPass := ""
	if config["db_pass"] != nil && len(config["db_pass"].(string)) > 0 {
		dbPass = config["db_pass"].(string)
	}

	return db.CreateMongoDbConnection(db.MongoConfig{
		uri, dbName, dbUser, dbPass,
	})
}
