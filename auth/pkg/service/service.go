package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthService describes the authentication service.
type AuthService interface {
	Authenticate(ctx context.Context, username, password string) (accessToken string, err error)
	ValidateToken(ctx context.Context, accessToken string) (userInfo *UserInfo, err error)
}

// UserInfo represents the user information.
type UserInfo struct {
	ID       string
	Username string
}

// auth implements the AuthService interface.
type auth struct {
	// You can add necessary dependencies or database connections here.
	dbService DbService // Replace with your actual db service interface
}

// NewAuthService creates a new instance of the authentication service.
func NewAuthService() AuthService {
	return &auth{}
}

// Authenticate handles the authentication request and generates an access token if the credentials are valid.
func (a *auth) Authenticate(ctx context.Context, username, password string) (accessToken string, err error) {
	// Retrieve the user from the database
	user, err := a.dbService.GetUserByUsername(ctx, username)
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
	accessToken, err = generateAccessToken(user.ID)
	if err != nil {
		return "", err
	}

	return accessToken, errors.New("invalid credentials")
}

// ValidateToken validates the access token and returns the user information.
func (a *auth) ValidateToken(ctx context.Context, accessToken string) (userInfo *UserInfo, err error) {
	if validateAccessToken(accessToken) {
		// Token is valid, retrieve user information from the token or database
		userID := extractUserIDFromToken(accessToken)
		user, err := a.dbService.GetUserByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		userInfo = &UserInfo{
			ID:       user.ID,
			Username: user.Username,
		}

		return userInfo, nil
	}

	return nil, errors.New("invalid access token")
}

// generateAccessToken generates a random access token.
func generateAccessToken(userID string) (string, error) {
	// Implement your own token generation logic.
	// You may use a library like github.com/dgrijalva/jwt-go or generate a random string.
	// Generate an access token using your preferred method or library
	// In this example, we will use a simple token format: userID + timestamp

	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Combine the userID and timestamp
	token := userID + "_" + string(timestamp)

	return token, nil
}

// validateAccessToken validates the access token against a database or blacklist.
func validateAccessToken(accessToken string) bool {
	// Implement your own token validation logic.
	// Verify the access token against a database of active tokens or blacklist.
	// Check the token's expiration and integrity.

	// Check if the token starts with "valid_"
	if !strings.HasPrefix(accessToken, "valid_") {
		return false
	}

	// Check the token expiration
	expiration := extractExpirationFromToken(accessToken)
	now := time.Now().Unix()

	if expiration <= now {
		return false
	}

	// Additional checks or validations can be performed here

	return true
}

func extractUserIDFromToken(accessToken string) string {
	// Implement your method to extract the user ID from the token
	// In this example, we assume the user ID is the substring after "valid_"
	// Modify this logic according to your actual token structure

	return strings.TrimPrefix(accessToken, "valid_")
}

func extractExpirationFromToken(accessToken string) int64 {
	// Implement your method to extract the expiration time from the token
	// In this example, we assume the expiration time is encoded as a UNIX timestamp
	// Modify this logic according to your actual token structure

	// Assuming the expiration timestamp is located after the prefix "valid_"
	expirationStr := strings.TrimPrefix(accessToken, "valid_")
	expiration, _ := strconv.ParseInt(expirationStr, 10, 64)

	return expiration
}