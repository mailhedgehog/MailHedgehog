package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailhedgehog/MailHedgehog/dto/smtpMessage"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/mail"
	"time"
)

type Mongo struct {
	Collection *mongo.Collection
}

func CreateMongoDbStorage(collection *mongo.Collection) *Mongo {
	indexName, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{"room", -1},
			{"id", 1},
		},
	})
	logger.PanicIfError(err)

	logManager().Debug(fmt.Sprintf("Index [%s] created", indexName))

	return &Mongo{
		Collection: collection,
	}
}

type Message struct {
	ID      smtpMessage.MessageID    `bson:"id"`
	Room    string                   `bson:"room"`
	From    []*mail.Address          `bson:"from"`
	To      []*mail.Address          `bson:"to"`
	Subject string                   `bson:"subject"`
	Date    time.Time                `bson:"date"`
	Message *smtpMessage.SMTPMessage `bson:"message"`
}

func (mongoClient *Mongo) RoomName(room Room) string {
	if len(room) <= 0 {
		room = "_default"
	}
	return room
}

func (mongoClient *Mongo) Store(room Room, message *smtpMessage.SMTPMail) (smtpMessage.MessageID, error) {
	if perRoomLimit > 0 && perRoomLimit <= mongoClient.Count(room) {
		mongoClient.DeleteRoom(room)
	}

	insertResult, err := mongoClient.Collection.InsertOne(context.TODO(), Message{
		message.ID,
		mongoClient.RoomName(room),
		message.Email.From,
		message.Email.To,
		message.Email.Subject,
		message.Email.Date,
		message.Origin,
	})

	logManager().Debug(fmt.Sprintf("New message saved, mongo _id='%s'", insertResult.InsertedID))

	return message.ID, err
}
func (mongoClient *Mongo) List(room Room, query SearchQuery, offset, limit int) ([]smtpMessage.SMTPMail, int, error) {

	opts := options.Find().SetSort(bson.M{"date": -1}).SetSkip(int64(offset)).SetLimit(int64(limit))

	textsMatch := bson.A{}
	for criteria, queryValue := range query {
		switch criteria {
		case "to":
			textsMatch = append(
				textsMatch,
				bson.M{"to.name": primitive.Regex{Pattern: queryValue, Options: ""}},
				bson.M{"to.address": primitive.Regex{Pattern: queryValue, Options: ""}},
			)
		case "from":
			textsMatch = append(
				textsMatch,
				bson.M{"from.name": primitive.Regex{Pattern: queryValue, Options: ""}},
				bson.M{"from.address": primitive.Regex{Pattern: queryValue, Options: ""}},
			)
		case "content":
			textsMatch = append(
				textsMatch,
				bson.M{"subject": primitive.Regex{Pattern: queryValue, Options: ""}},
			)
		}
	}

	filterQuery := bson.A{
		bson.M{"room": mongoClient.RoomName(room)},
	}
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

	var emailsList []smtpMessage.SMTPMail
	var results []Message
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, 0, err
	}
	for _, result := range results {
		m, err := result.Message.ToSMTPMail(result.ID)
		if err != nil {
			return nil, 0, err
		}
		emailsList = append(emailsList, *m)
	}

	return emailsList, int(totalCount), nil
}
func (mongoClient *Mongo) Count(room Room) int {
	count, err := mongoClient.Collection.CountDocuments(context.TODO(), bson.D{{"room", mongoClient.RoomName(room)}})
	logger.PanicIfError(err)

	return int(count)
}
func (mongoClient *Mongo) Delete(room Room, messageId smtpMessage.MessageID) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"room", mongoClient.RoomName(room)}},
				bson.D{{"id", messageId}},
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
func (mongoClient *Mongo) Load(room Room, messageId smtpMessage.MessageID) (*smtpMessage.SMTPMail, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"room", mongoClient.RoomName(room)}},
				bson.D{{"id", messageId}},
			}},
	}
	var result Message
	err := mongoClient.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	m, err := result.Message.ToSMTPMail(messageId)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mongoClient *Mongo) RoomsList(offset, limit int) ([]Room, error) {
	cursor, err := mongoClient.Collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$group", bson.M{"_id": "$room"}}},
		bson.D{{"$sort", bson.M{"_id": 1}}},
		bson.D{{"$skip", offset}},
		bson.D{{"$limit", limit}},
	},
	)
	logger.PanicIfError(err)

	var results []Room
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		logger.PanicIfError(err)
		results = append(results, result["_id"].(string))
	}

	logger.PanicIfError(cursor.Err())

	return results, nil
}

func (mongoClient *Mongo) RoomsCount() int {
	cursor, err := mongoClient.Collection.Aggregate(context.TODO(), mongo.Pipeline{bson.D{
		{"$group", bson.D{
			{"_id", "$room"},
			{"count", bson.M{"$sum": 1}},
		}}}},
	)
	logger.PanicIfError(err)

	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	logger.PanicIfError(err)
	if len(results) > 0 {
		return int(results[0]["count"].(int32))
	}

	return 0
}

func (mongoClient *Mongo) DeleteRoom(room Room) error {
	roomName := mongoClient.RoomName(room)
	result, err := mongoClient.Collection.DeleteMany(context.TODO(), bson.M{"room": roomName})

	logManager().Debug(fmt.Sprintf("Deleted room [%s] (%d items)", roomName, result.DeletedCount))

	return err
}
