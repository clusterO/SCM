package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	driverbson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

/* Business logic */

// DbService describes the database service.
type DbService interface {
	SaveUser(user *User) error
	GetUserByID(userID string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type dbService struct{}

// User represents a user in the system.
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

// SaveUser saves a user to the database.
func (dbService) SaveUser(user *User) error {
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
		return errors.New("user already exists")
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID retrieves a user from the database based on the user ID.
func (dbService) GetUserByID(userID string) (*User, error) {
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
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrieves a user from the database based on the username.
func (dbService) GetUserByUsername(username string) (*User, error) {
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
		return nil, err
	}


	return &user, nil
}

// getDBName retrieves the database name from the environment variables.
func getDBName() string {
	return os.Getenv("DB_NAME")
}

// getDBUrl retrieves the database url from the environment variables.
func getDBUrl() string {
	return os.Getenv("DB_URL")
}

// loadEnv loads the environment variables from the .env file.
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

// Initialize the environment variables.
func init() {
	err := loadEnv()
	if err != nil {
		panic("Failed to load .env file")
	}
}

/* Requests and responses */

// SaveUserRequest represents the request parameters for the SaveUser method.
type SaveUserRequest struct {
	User *User
}

// SaveUserResponse represents the response for the SaveUser method.
type SaveUserResponse struct {
	Err error
}

// GetUserByIDRequest represents the request parameters for the GetUserByID method.
type GetUserByIDRequest struct {
	UserID string
}

// GetUserByIDResponse represents the response for the GetUserByID method.
type GetUserByIDResponse struct {
	User *User
	Err  error
}

// GetUserByUsernameRequest represents the request parameters for the GetUserByUsername method.
type GetUserByUsernameRequest struct {
	Username string
}

// GetUserByUsernameResponse represents the response for the GetUserByUsername method.
type GetUserByUsernameResponse struct {
	User *User
	Err  error
}

/* Endpoints */

func MakeSaveUserEndpoint(s DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*SaveUserRequest)
		err := s.SaveUser(req.User)
		return &SaveUserResponse{Err: err}, nil
	}
}

func MakeGetUserByIDEndpoint(s DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetUserByIDRequest)
		user, err := s.GetUserByID(req.UserID)
		return &GetUserByIDResponse{User: user, Err: err}, nil
	}
}

func MakeGetUserByUsernameEndpoint(s DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetUserByUsernameRequest)
		user, err := s.GetUserByUsername(req.Username)
		return &GetUserByUsernameResponse{User: user, Err: err}, nil
	}
}

/* Transports */

func main() {
	dbs := dbService{}

	SaveUserHandler := httptransport.NewServer(
		MakeSaveUserEndpoint(dbs),
		decodeSaveUserRequest,
		encodeResponse,
	)

	GetUserByIDHandler := httptransport.NewServer(
		MakeGetUserByIDEndpoint(dbs),
		decodeGetUserByIDRequest,
		encodeResponse,
	)

	GetUserByUsernameHandler := httptransport.NewServer(
		MakeGetUserByUsernameEndpoint(dbs),
		decodeGetUserByUsernameRequest,
		encodeResponse,
	)

	http.Handle("/save", SaveUserHandler)
	http.Handle("/get_by_id", GetUserByIDHandler)
	http.Handle("/get_by_username", GetUserByUsernameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// decodeSaveUserRequest decodes the incoming HTTP request into a SaveUserRequest struct.
func decodeSaveUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SaveUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserByIDRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeGetUserByUsernameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserByUsernameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
