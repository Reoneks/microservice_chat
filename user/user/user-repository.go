package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	usersCollection *mongo.Collection
	ctx             context.Context
}

type IUserRepository interface {
	GetUserByID(id string) (map[string]interface{}, error)
	GetUsers() ([]map[string]interface{}, error)
	CreateUser(user map[string]interface{}) (map[string]interface{}, error)
	UpdateUser(user map[string]interface{}) (map[string]interface{}, error)
	DeleteUser(id string) error
}

func (ur *userRepository) GetUserByID(id string) (map[string]interface{}, error) {
	filter := bson.M{"_id": id}

	var user map[string]interface{}
	err := ur.usersCollection.FindOne(ur.ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("FindOne error:\n\t%v", err)
	}

	return user, nil
}

func (ur *userRepository) GetUsers() ([]map[string]interface{}, error) {
	cur, err := ur.usersCollection.Find(ur.ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("Find error:\n\t%v", err)
	}

	defer cur.Close(ur.ctx)

	var users []map[string]interface{}
	for cur.Next(ur.ctx) {
		var data map[string]interface{}

		err := cur.Decode(&data)
		if err != nil {
			log.Println(fmt.Errorf("Decode error:\n\t%v", err))
			continue
		}

		users = append(users, data)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("Cursor error:\n\t%v", err)
	}

	return users, nil
}

func (ur *userRepository) CreateUser(user map[string]interface{}) (map[string]interface{}, error) {
	user["_id"] = uuid.New().String()
	user["created_at"] = time.Now()
	user["updated_at"] = time.Now()

	_, err := ur.usersCollection.InsertOne(ur.ctx, user)
	if err != nil {
		return nil, fmt.Errorf("CreateUser error:\n\t%v", err)
	}

	return user, nil
}

func (ur *userRepository) UpdateUser(user map[string]interface{}) (map[string]interface{}, error) {
	user["updated_at"] = time.Now()

	filter := bson.M{"_id": user["_id"]}
	update := bson.M{"$set": user}
	_, err := ur.usersCollection.UpdateOne(ur.ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOne error:\n\t%v", err)
	}

	return user, nil
}

func (ur *userRepository) DeleteUser(id string) error {
	_, err := ur.usersCollection.DeleteOne(ur.ctx, bson.M{"_id": id})
	return err
}

func NewUserRepository(db *mongo.Client, dbName, collection string) IUserRepository {
	usersCollection := db.Database(dbName).Collection(collection)
	return &userRepository{usersCollection, context.Background()}
}
