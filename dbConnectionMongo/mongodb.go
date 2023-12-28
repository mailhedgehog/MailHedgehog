package dbConnectionMongo

import (
	"context"
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ConnectionConfig struct {
	URI  string
	DB   string
	User string
	Pass string
}

func MakeCollection(config contracts.DbConnectionConfig, collection string) *mongo.Collection {
	return MakeConnection(MakeConfigurationFromContract(config)).Collection(collection)
}

func MakeConfigurationFromContract(config contracts.DbConnectionConfig) *ConnectionConfig {
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

	return &ConnectionConfig{
		uri, dbName, dbUser, dbPass,
	}
}

func MakeConnection(config *ConnectionConfig) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.URI).SetTimeout(5 * time.Second)

	if len(config.User) > 0 {
		clientOptions = clientOptions.SetAuth(options.Credential{
			Username: config.User,
			Password: config.Pass,
		})
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	logger.PanicIfError(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	logger.PanicIfError(err)

	// logManager().Debug("Connected to MongoDB")

	return client.Database(config.DB)
}
