package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	db "db/pkg/service"
)

// AuthService describes the authentication service.
type AuthService interface {
	Authenticate(ctx context.Context, username, password string) (accessToken string, err error)
	Authorize(ctx context.Context, token string, permission string) (bool, error) // Separate role management to its own service
	ValidateToken(ctx context.Context, accessToken string) (userInfo *(db.User), err error)
	Encryption(ctx context.Context, data []byte) ([]byte, error)
	// OAuth2 for social accounts connection
}

// auth implements the AuthService interface.
type Auth struct {
	// You can add necessary dependencies or database connections here.
	dbService db.DbService // Replace with your actual db service interface
}

// Authenticate handles the authentication request and generates an access token if the credentials are valid.
func (a Auth) Authenticate(ctx context.Context, username, password string) (accessToken string, err error) {
	user, err := a.dbService.GetUserByUsername(username) // a.dbService is nil FIX
	if err != nil {
		return "", err
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Password does not match
		return "", errors.New("invalid credentials")
	}

	// Generate an access token
	accessToken, err = generateAccessToken(string(user.ID))
	if err != nil {
		return "", err
	}

	return accessToken, errors.New("invalid credentials")
}

func (a Auth) Authorize(ctx context.Context, accessToken string, permission string) (bool, error) {
	userInfo, err := a.ValidateToken(ctx, accessToken)
	if err != nil {
		return false, err
	}

	// Check if the user has the required permission
	hasPermission := checkUserPermission(userInfo, permission)
	return hasPermission, nil
}

// ValidateToken validates the access token and returns the user information.
func (a Auth) ValidateToken(ctx context.Context, accessToken string) (userInfo *(db.User), err error) {
	if validateAccessToken(accessToken) {
		// Token is valid, retrieve user information from the token or database
		userID := extractUserIDFromToken(accessToken)
		user, err := a.dbService.GetUserByID(userID)
		if err != nil {
			return nil, err
		}

		userInfo = &db.User {
			ID:       user.ID,
			Username: user.Username,
		}

		return userInfo, nil
	}

	return nil, errors.New("invalid access token")
}

func (a Auth) Encryption(ctx context.Context, data []byte) ([]byte, error) {
	// SSH Encryption
	sshPublicKey, err := generateSSHPublicKey()
	if err != nil {
		return nil, err
	}

	encryptedData, err := encryptWithSSH(data, sshPublicKey)
	if err != nil {
		return nil, err
	}

	// SSL Encryption
	encryptedData, err = encryptWithSSL(encryptedData)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}