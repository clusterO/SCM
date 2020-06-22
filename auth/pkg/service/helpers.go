package auth

import (
	db "db/pkg/service"
	"strconv"
	"strings"
	"time"
)

func checkUserPermission(userInfo *(db.User), permission string) bool {
	for _, role := range userInfo.Roles {
		if hasRolePermission(role, permission) {
			return true
		}
	}

	return false
}

func hasRolePermission(role string, permission string) bool {
	// You may check against a database, configuration file, or other rules to validate the role permission.
	if role == "admin" && permission == "admin" {
		return true
	}

	return false
}

// Enhance

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