package auth

import (
	"context"
	db "db/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

/* Requests and responses */

// AuthenticateRequest struct
type AuthenticateRequest struct {
	Username string
	Password string
}

// AuthenticateResponse struct
type AuthenticateResponse struct {
	AccessToken string
	Error       error
}

// AuthorizeRequest struct
type AuthorizeRequest struct {
	Token      string
	Permission string
}

// AuthorizeResponse struct
type AuthorizeResponse struct {
	Authorized bool
	Error      error
}

// ValidateTokenRequest struct
type ValidateTokenRequest struct {
	AccessToken string
}

// ValidateTokenResponse struct
type ValidateTokenResponse struct {
	UserInfo *(db.User)
	Error    error
}

// EncryptionRequest struct
type EncryptionRequest struct {
	Data []byte
}

// EncryptionResponse struct
type EncryptionResponse struct {
	EncryptedData []byte
	Error         error
}

/* Endpoints */

// Authenticate endpoint
func MakeAuthenticateEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthenticateRequest)
		accessToken, err := svc.Authenticate(ctx, req.Username, req.Password)
		return &AuthenticateResponse{AccessToken: accessToken, Error: err}, nil
	}
}

// Authorize endpoint
func MakeAuthorizeEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthorizeRequest)
		authorized, err := svc.Authorize(ctx, req.Token, req.Permission)
		return &AuthorizeResponse{Authorized: authorized, Error: err}, nil
	}
}

// ValidateToken endpoint
func MakeValidateTokenEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateTokenRequest)
		userInfo, err := svc.ValidateToken(ctx, req.AccessToken)
		return &ValidateTokenResponse{UserInfo: userInfo, Error: err}, nil
	}
}

// Encryption endpoint
func MakeEncryptionEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EncryptionRequest)
		encryptedData, err := svc.Encryption(ctx, req.Data)
		return &EncryptionResponse{EncryptedData: encryptedData, Error: err}, nil
	}
}

/* Transports -- publicly accessible FIX */

// DecodeAuthenticateRequest function
func DecodeAuthenticateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AuthenticateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeAuthorizeRequest function
func DecodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AuthorizeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// DecodeValidateTokenRequest function
func DecodeValidateTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ValidateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// DecodeEncryptionRequest function
func DecodeEncryptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req EncryptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// EncodeResponse function
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// ErrorEncoder function
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}