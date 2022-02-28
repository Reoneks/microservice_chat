package messages

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessagesRepository interface {
	GetMessagesByRoom(roomID string, limit, offset int64) ([]map[string]interface{}, error)
	GetMessageByID(id string) (map[string]interface{}, error)
	CreateMessage(message map[string]interface{}) (map[string]interface{}, error)
	UpdateMessage(message map[string]interface{}) (map[string]interface{}, error)
	DeleteMessage(id string) error
}

type MessagesRepositoryImpl struct {
	messagesCollection *mongo.Collection
	ctx                context.Context
}

func (r *MessagesRepositoryImpl) GetMessagesByRoom(roomID string, limit, offset int64) ([]map[string]interface{}, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit).SetSkip(offset).SetSort(bson.M{"updated_at": -1})

	cur, err := r.messagesCollection.Find(r.ctx, bson.M{"room_id": roomID}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("Find error:\n\t%v", err)
	}

	defer cur.Close(r.ctx)

	var results []map[string]interface{}
	for cur.Next(r.ctx) {
		var data map[string]interface{}

		err := cur.Decode(&data)
		if err != nil {
			log.Println(fmt.Errorf("Decode error:\n\t%v", err))
			continue
		}

		results = append(results, data)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("Cursor error:\n\t%v", err)
	}

	return results, nil
}

func (r *MessagesRepositoryImpl) GetMessageByID(id string) (map[string]interface{}, error) {
	message := make(bson.M)
	err := r.messagesCollection.FindOne(r.ctx, bson.M{"_id": id}).Decode(&message)
	if err != nil {
		return nil, fmt.Errorf("CreateMessage error:\n\t%v", err)
	}

	return message, nil
}

func (r *MessagesRepositoryImpl) CreateMessage(message map[string]interface{}) (map[string]interface{}, error) {
	message["_id"] = uuid.New().String()
	message["created_at"] = time.Now()
	message["updated_at"] = time.Now()
	_, err := r.messagesCollection.InsertOne(r.ctx, message)
	if err != nil {
		return nil, fmt.Errorf("CreateMessage error:\n\t%v", err)
	}

	return message, nil
}

func (r *MessagesRepositoryImpl) UpdateMessage(message map[string]interface{}) (map[string]interface{}, error) {
	message["updated_at"] = time.Now()

	filter := bson.M{"_id": message["_id"]}
	update := bson.M{"$set": message}
	_, err := r.messagesCollection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOne error:\n\t%v", err)
	}
	return message, nil
}

func (r *MessagesRepositoryImpl) DeleteMessage(id string) error {
	_, err := r.messagesCollection.DeleteOne(r.ctx, bson.M{"_id": id})
	return err
}

func NewMessagesRepository(db *mongo.Client, dbName, collection string) MessagesRepository {
	messagesCollection := db.Database(dbName).Collection(collection)
	return &MessagesRepositoryImpl{messagesCollection, context.Background()}
}
