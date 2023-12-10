package db

import (
	"context"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConfig struct {
	URI  string
	DB   string
	User string
	Pass string
}

func CreateMongoDbConnection(config MongoConfig) *mongo.Database {
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

	logManager().Debug("Connected to MongoDB")

	return client.Database(config.DB)
}

func CreateMongoDbCollectionConnection(config MongoConfig, collection string) *mongo.Collection {
	return CreateMongoDbConnection(config).Collection(collection)
}
