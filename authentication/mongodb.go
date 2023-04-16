package authentication

import (
	"context"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/db"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Mongo struct {
	Collection *mongo.Collection
}

type UserRow struct {
	Username string `bson:"username"`
	HttpPass string `bson:"http_password"`
	SmtpPass string `bson:"smtp_password"`
}

func CreateMongoDbAuthentication(config db.MongoConfig) *Mongo {
	collection := db.CreateMongoDbCollectionConnection(config)

	indexName, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.M{"username": 1},
	})
	logger.PanicIfError(err)

	logManager().Debug(fmt.Sprintf("Index [%s] created", indexName))

	return &Mongo{
		Collection: collection,
	}
}

func (mongoClient *Mongo) RequiresAuthentication() bool {
	return true
}

func (mongoClient *Mongo) Authenticate(authType AuthenticationType, username string, password string) bool {
	if !mongoClient.RequiresAuthentication() {
		return true
	}

	var user UserRow
	err := mongoClient.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	passwordHashToCheck := user.HttpPass
	if authType == SMTP && len(user.SmtpPass) > 0 {
		passwordHashToCheck = user.SmtpPass
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHashToCheck), []byte(password)); err != nil {
		return false
	}

	return true
}

func (mongoClient *Mongo) UsernamePresent(username string) bool {
	count, err := mongoClient.Collection.CountDocuments(context.TODO(), bson.M{"username": username})

	if err != nil || count <= 0 {
		return false
	}

	return true
}

func (mongoClient *Mongo) AddUser(username string, httpPassHash string, smtpPassHash string) error {
	insertResult, err := mongoClient.Collection.InsertOne(context.TODO(), UserRow{
		username,
		httpPassHash,
		smtpPassHash,
	})

	fmt.Println(username, httpPassHash, smtpPassHash)

	logManager().Debug(fmt.Sprintf("New useer [%s] added, mongo _id='%s'", username, insertResult.InsertedID))

	return err
}
