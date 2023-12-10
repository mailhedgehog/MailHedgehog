package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

type Mongo struct {
	Collection *mongo.Collection
}

type UserRow struct {
	Username      string   `bson:"username"`
	HttpPass      string   `bson:"http_password"`
	SmtpPass      string   `bson:"smtp_password"`
	NoPassIPs     []string `bson:"no_pass_ips"`
	RestrictedIPs []string `bson:"restricted_ips"`
	LoginEmails   []string `bson:"login_emails"`
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

func (mongoClient *Mongo) AuthenticateSMTPViaIP(username string, ip string) bool {
	if !mongoClient.RequiresAuthentication() {
		return true
	}

	var user UserRow
	err := mongoClient.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	return slices.Contains(user.NoPassIPs, ip)
}

func (mongoClient *Mongo) SmtpIpIsWhitelisted(username string, ip string) bool {
	if !mongoClient.RequiresAuthentication() {
		return true
	}

	var user UserRow
	err := mongoClient.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	if len(user.RestrictedIPs) > 0 {
		return slices.Contains(user.RestrictedIPs, ip)
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
		[]string{},
		[]string{},
		[]string{},
	})

	logManager().Debug(fmt.Sprintf("New user [%s] added, mongo _id='%s'", username, insertResult.InsertedID))

	return err
}

func (mongoClient *Mongo) UpdateUser(username string, httpPassHash string, smtpPassHash string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	if len(httpPassHash) > 0 {
		newValues = append(newValues, bson.E{"http_password", httpPassHash})
	}
	if len(smtpPassHash) > 0 {
		newValues = append(newValues, bson.E{"smtp_password", smtpPassHash})
	}

	updateResult, err := mongoClient.Collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

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
		resources = append(resources, UserResource{
			Username:      result.Username,
			NoPassIPs:     result.NoPassIPs,
			RestrictedIPs: result.RestrictedIPs,
			LoginEmails:   result.LoginEmails,
		})
	}

	return resources, int(totalCount), nil
}

func (mongoClient *Mongo) AddNoPassSmtpIp(username string, ip string) error {
	var user UserRow
	err := mongoClient.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	if slices.Contains(user.NoPassIPs, ip) {
		return nil
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	user.NoPassIPs = append(user.NoPassIPs, ip)

	newValues = append(newValues, bson.E{"NoPassIPs", user.NoPassIPs})

	updateResult, err := mongoClient.Collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (mongoClient *Mongo) DeleteNoPassSmtpIp(username string, ip string) error {
	var user UserRow
	err := mongoClient.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	if !slices.Contains(user.NoPassIPs, ip) {
		return nil
	}

	i := slices.Index(user.NoPassIPs, ip)
	user.NoPassIPs = slices.Delete(user.NoPassIPs, i, i+1)

	newValues = append(newValues, bson.E{"NoPassIPs", user.NoPassIPs})

	updateResult, err := mongoClient.Collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (mongoClient *Mongo) ClearAllNoPassSmtpIps(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	newValues = append(newValues, bson.E{"NoPassIPs", []string{}})

	updateResult, err := mongoClient.Collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}
