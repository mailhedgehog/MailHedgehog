package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func CreateMongoDbAuthentication(collection *mongo.Collection) *Mongo {
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

func (mongoClient *Mongo) DeleteUser(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	result, err := mongoClient.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New(fmt.Sprintf("Unexpected count of deleted items, extected 1, got %d", result.DeletedCount))
	}

	return nil
}

func (mongoClient *Mongo) ListUsers(searchQuery string, offset, limit int) ([]UserResource, int, error) {
	opts := options.Find().SetSort(bson.M{"username": 1}).SetSkip(int64(offset)).SetLimit(int64(limit))
	textsMatch := bson.A{
		bson.M{"username": primitive.Regex{Pattern: searchQuery, Options: ""}},
	}
	filterQuery := bson.A{}
	if len(textsMatch) > 0 {
		filterQuery = append(filterQuery, bson.M{"$or": textsMatch})
	}
	filter := bson.M{"$and": filterQuery}

	totalCount, err := mongoClient.Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := mongoClient.Collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, 0, err
	}
	var resources []UserResource
	var results []UserRow
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, 0, err
	}
	for _, result := range results {
		resources = append(resources, UserResource{Username: result.Username})
	}

	return resources, int(totalCount), nil
}
