package dbservice

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* Business logic */

// DbService describes the database service.
type DbService interface {
	SaveUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type DBService struct{}

// User represents a user in the system.
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

// db implements the DbService interface.
type db struct {
	session *mgo.Session
}

// NewDbService creates a new instance of the database service.
func NewDbService(session *mgo.Session) DbService {
	return &db{session: session}
}

// SaveUser saves a user to the database.
func (d *db) SaveUser(ctx context.Context, user *User) error {
	session := d.session.Copy()
	defer session.Close()

	collection := session.DB(getDBName()).C("users")

	// Check if the user already exists in the database
	existingUser := User{}
	err := collection.Find(bson.M{"username": user.Username}).One(&existingUser)
	if err == nil {
		return errors.New("user already exists")
	}

	err = collection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID retrieves a user from the database based on the user ID.
func (d *db) GetUserByID(ctx context.Context, userID string) (*User, error) {
	session := d.session.Copy()
	defer session.Close()

	collection := session.DB(getDBName()).C("users")

	user := User{}
	err := collection.FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrieves a user from the database based on the username.
func (d *db) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	session := d.session.Copy()
	defer session.Close()

	collection := session.DB(getDBName()).C("users")

	user := User{}
	err := collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// getDBName retrieves the database name from the environment variables.
func getDBName() string {
	return os.Getenv("DB_NAME")
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

/* Endpoints */

type Endpoints struct {
	SaveUserEndpoint          endpoint.Endpoint
	GetUserByIDEndpoint       endpoint.Endpoint
	GetUserByUsernameEndpoint endpoint.Endpoint
}

func NewEndpoints(s DbService) Endpoints {
	return Endpoints{
		SaveUserEndpoint:          MakeSaveUserEndpoint(s),
		GetUserByIDEndpoint:       MakeGetUserByIDEndpoint(s),
		GetUserByUsernameEndpoint: MakeGetUserByUsernameEndpoint(s),
	}
}

/* Transports */

func NewHTTPHandler(endpoints Endpoints) http.Handler {
	mux := http.NewServeMux()

	options := []httptransport.ServerOption{}

	mux.Handle("/saveUser", httptransport.NewServer(
		endpoints.SaveUserEndpoint,
		decodeSaveUserRequest,
		encodeResponse,
		options...,
	))

	mux.Handle("/getUserByID", httptransport.NewServer(
		endpoints.GetUserByIDEndpoint,
		decodeGetUserByIDRequest,
		encodeResponse,
		options...,
	))

	mux.Handle("/getUserByUsername", httptransport.NewServer(
		endpoints.GetUserByUsernameEndpoint,
		decodeGetUserByUsernameRequest,
		encodeResponse,
		options...,
	))

	return mux
}
