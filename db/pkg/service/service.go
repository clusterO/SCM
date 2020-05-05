package db

import (
	"context"
	"errors"

	driverbson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbService describes the service.
type DbService interface {
	SaveUser(user *User) error
	GetUserByID(userID string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

/* Business logic */
// -- publicly accessible FIX
type DBService struct{}


/* type DBService struct{
	repository Repository
	logger log.Logger
}*/

/*
func NewDbService(rep Repository, logger log.Logger) DbService {
	return &dbService{
		repository: rep,
		logger: logger,
	}
}
*/

// SaveUser saves a user to the database.
func (dbs DBService) SaveUser(user *User) error {
	// logger := log.With(dbs.logger, "method", "SaveUser")

	/*** TO DO through repository ***/
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getDBUrl()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database(getDBName()).Collection("users")

	// Check if the user already exists in the database
	existingUser := User{}
	err = collection.FindOne(context.TODO(), driverbson.D{{Key: "username", Value: user.Username}}).Decode(&existingUser)
	if err == nil {
		// level.Error(logger).Log("err", err)
		return errors.New("User already exists")
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	// logger.Log("User created", id)

	return nil
}

// GetUserByID retrieves a user from the database based on the user ID.
func (dbs DBService) GetUserByID(userID string) (*User, error) {
	//logger := log.With(dbs.logger, "method", "GetUserByID")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getDBUrl()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database(getDBName()).Collection("users")

	user := User{}
	err = collection.FindOne(context.TODO(), driverbson.D{{Key: "_id", Value: userID}}).Decode(&user)
	if err != nil {
		// level.Error(logger).Log("err", err)
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrieves a user from the database based on the username.
func (dbs DBService) GetUserByUsername(username string) (*User, error) {
	// logger := log.With(dbs.logger, "method", "GetUserByUsername")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getDBUrl()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	collection := client.Database(getDBName()).Collection("users")

	user := User{}
	err = collection.FindOne(context.TODO(), driverbson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		// level.Error(logger).Log("err", err)
		return nil, err
	}

	return &user, nil
}
