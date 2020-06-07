// Move to folder "models"
package db

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

// User represents a user in the system.
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Roles    []string
}

// Architecture enhance
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}