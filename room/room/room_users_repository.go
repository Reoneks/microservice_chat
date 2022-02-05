package room

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomUsersRepository interface {
	GetRoomsByUserId(id string, limit, offset int64) ([]map[string]interface{}, error)
	CreateUserRoomConnection(user map[string]interface{}) (map[string]interface{}, error)
	DeleteUserRoomConnection(user map[string]interface{}) error
}

type RoomUsersRepositoryImpl struct {
	roomsUsersCollection *mongo.Collection
	roomsCollection      *mongo.Collection
	ctx                  context.Context
}

func (r *RoomUsersRepositoryImpl) GetRoomsByUserId(id string, limit, offset int64) ([]map[string]interface{}, error) {
	var rooms []map[string]interface{}

	cur, err := r.roomsUsersCollection.Find(r.ctx, bson.M{"user_id": id})
	if err != nil {
		return nil, fmt.Errorf("Find error:\n\t%v", err)
	}

	defer cur.Close(r.ctx)

	var roomUsers []map[string]interface{}
	for cur.Next(r.ctx) {
		var data map[string]interface{}

		err := cur.Decode(&data)
		if err != nil {
			log.Println(fmt.Errorf("Decode error:\n\t%v", err))
			continue
		}

		roomUsers = append(roomUsers, data)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("Cursor error:\n\t%v", err)
	}

	var IDs []interface{}
	for _, user := range roomUsers {
		IDs = append(IDs, user["room_id"])
	}

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(offset)

	rcur, err := r.roomsCollection.Find(r.ctx, bson.M{"_id": bson.M{"$in": IDs}}, opts)
	if err != nil {
		return nil, fmt.Errorf("Find error:\n\t%v", err)
	}

	defer rcur.Close(r.ctx)

	for rcur.Next(r.ctx) {
		var data map[string]interface{}

		err := rcur.Decode(&data)
		if err != nil {
			log.Println(fmt.Errorf("Decode error:\n\t%v", err))
			continue
		}

		rooms = append(rooms, data)
	}

	if err := rcur.Err(); err != nil {
		return nil, fmt.Errorf("Cursor error:\n\t%v", err)
	}

	return rooms, nil
}

func (r *RoomUsersRepositoryImpl) CreateUserRoomConnection(user map[string]interface{}) (map[string]interface{}, error) {
	_, err := r.roomsUsersCollection.InsertOne(r.ctx, user)
	if err != nil {
		return nil, fmt.Errorf("CreateUserRoomConnection error:\n\t%v", err)
	}
	return user, nil
}

func (r *RoomUsersRepositoryImpl) DeleteUserRoomConnection(user map[string]interface{}) error {
	_, err := r.roomsUsersCollection.DeleteOne(r.ctx, user)
	return err
}

func NewRoomUsersRepository(db *mongo.Client, dbName, roomUserCollection, roomCollection string) RoomUsersRepository {
	roomsUsersCollection := db.Database(dbName).Collection(roomUserCollection)
	roomsCollection := db.Database(dbName).Collection(roomCollection)
	return &RoomUsersRepositoryImpl{roomsUsersCollection, roomsCollection, context.Background()}
}
