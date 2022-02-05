package auth

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	GetUserByEmail(email string) (authUser map[string]interface{}, err error)
	AddUser(authUser map[string]interface{}) (map[string]interface{}, error)
	DeleteUser(id string) error
}

type AuthRepositoryImpl struct {
	authCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthRepository(db *mongo.Client, dbName, collection string) AuthRepository {
	authCollection := db.Database(dbName).Collection(collection)
	return &AuthRepositoryImpl{authCollection, context.Background()}
}

func (a *AuthRepositoryImpl) GetUserByEmail(email string) (map[string]interface{}, error) {
	filter := bson.M{"email": email}

	var auth map[string]interface{}
	err := a.authCollection.FindOne(a.ctx, filter).Decode(&auth)
	if err != nil {
		return nil, fmt.Errorf("FindOne error:\n\t%v", err)
	}

	return auth, nil
}

func (a *AuthRepositoryImpl) AddUser(authUser map[string]interface{}) (map[string]interface{}, error) {
	_, err := a.authCollection.InsertOne(a.ctx, authUser)
	if err != nil {
		return nil, fmt.Errorf("AddUser error:\n\t%v", err)
	}

	return authUser, nil
}

func (a *AuthRepositoryImpl) DeleteUser(id string) error {
	_, err := a.authCollection.DeleteOne(a.ctx, bson.M{"_id": id})
	return err
}
