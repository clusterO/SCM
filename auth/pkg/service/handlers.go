package auth

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

// use NewHandler and make the functions private !!
	
func AuthenticateHandler(oths AuthService) *httptransport.Server { 
	return httptransport.NewServer(
	MakeAuthenticateEndpoint(oths),
	DecodeAuthenticateRequest,
	EncodeResponse,
	httptransport.ServerErrorEncoder(ErrorEncoder),
)}

func AuthorizeHandler(oths AuthService) *httptransport.Server {
	return httptransport.NewServer(
	MakeAuthorizeEndpoint(oths),
	DecodeAuthorizeRequest,
	EncodeResponse,
	httptransport.ServerErrorEncoder(ErrorEncoder),
)}

func ValidateTokenHandler(oths AuthService) *httptransport.Server { 
	return httptransport.NewServer(
	MakeValidateTokenEndpoint(oths),
	DecodeValidateTokenRequest,
	EncodeResponse,
	httptransport.ServerErrorEncoder(ErrorEncoder),
)}

func EncryptionHandler(oths AuthService) *httptransport.Server {
	return httptransport.NewServer(
	MakeEncryptionEndpoint(oths),
	DecodeEncryptionRequest,
	EncodeResponse,
	httptransport.ServerErrorEncoder(ErrorEncoder),
)}
