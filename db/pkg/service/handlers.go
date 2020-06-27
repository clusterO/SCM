package db

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

func SaveUserHandler(dbs DbService) *httptransport.Server { 
	return httptransport.NewServer(
	MakeSaveUserEndpoint(dbs),
	DecodeSaveUserRequest,
	EncodeResponse,
)}

func GetUserByIDHandler(dbs DbService) *httptransport.Server { 
	return httptransport.NewServer(
	MakeGetUserByIDEndpoint(dbs),
	DecodeGetUserByIDRequest,
	EncodeResponse,
)}

func GetUserByUsernameHandler(dbs DbService) *httptransport.Server {
	return httptransport.NewServer(
	MakeGetUserByUsernameEndpoint(dbs),
	DecodeGetUserByUsernameRequest,
	EncodeResponse,
)}