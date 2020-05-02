package dbservice

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

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
		err := s.SaveUser(ctx, req.User)
		return &SaveUserResponse{Err: err}, nil
	}
}

func MakeGetUserByIDEndpoint(s DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetUserByIDRequest)
		user, err := s.GetUserByID(ctx, req.UserID)
		return &GetUserByIDResponse{User: user, Err: err}, nil
	}
}

func MakeGetUserByUsernameEndpoint(s DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetUserByUsernameRequest)
		user, err := s.GetUserByUsername(ctx, req.Username)
		return &GetUserByUsernameResponse{User: user, Err: err}, nil
	}
}

/* Transports */

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
