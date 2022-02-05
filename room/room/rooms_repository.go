package room

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type roomRepository struct {
	roomsCollection *mongo.Collection
	ctx             context.Context
}

type IRoomRepository interface {
	GetRoom(id string) (map[string]interface{}, error)
	CreateRoom(user map[string]interface{}) (map[string]interface{}, error)
	UpdateRoom(user map[string]interface{}) (map[string]interface{}, error)
	DeleteRoom(id string) error
}

func (r *roomRepository) GetRoom(id string) (map[string]interface{}, error) {
	filter := bson.M{"_id": id}

	var room map[string]interface{}
	err := r.roomsCollection.FindOne(r.ctx, filter).Decode(&room)
	if err != nil {
		return nil, fmt.Errorf("FindOne error:\n\t%v", err)
	}

	return room, nil
}

func (r *roomRepository) CreateRoom(room map[string]interface{}) (map[string]interface{}, error) {
	room["_id"] = uuid.New().String()
	room["created_at"] = time.Now()
	room["updated_at"] = time.Now()

	_, err := r.roomsCollection.InsertOne(r.ctx, room)
	if err != nil {
		return nil, fmt.Errorf("CreateRoom error:\n\t%v", err)
	}

	return room, nil
}

func (r *roomRepository) UpdateRoom(room map[string]interface{}) (map[string]interface{}, error) {
	room["updated_at"] = time.Now()

	filter := bson.M{"_id": room["_id"]}
	update := bson.M{"$set": room}
	_, err := r.roomsCollection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOne error:\n\t%v", err)
	}

	return room, nil
}

func (r *roomRepository) DeleteRoom(id string) error {
	_, err := r.roomsCollection.DeleteOne(r.ctx, bson.M{"_id": id})
	return err
}

func NewRoomRepository(db *mongo.Client, dbName, collection string) IRoomRepository {
	roomsCollection := db.Database(dbName).Collection(collection)
	return &roomRepository{roomsCollection, context.Background()}
}
