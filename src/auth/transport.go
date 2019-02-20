package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeAuthorizeEndpoint(svc AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(authorizeRequest)
		//Validate req.Username & req.Password
		auth := svc.Authorize()
		//if err != nil {
		//	return authroizeRequest{r, err.Error()}, nil
		//}
		return authorizeResponse{auth, ""}, nil
	}
}

func MakeAccessTokenEndpoint(svc AuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(accesstokenRequest)
		token := svc.AccessToken()
		//if err != nil {
		//	return authroizeRequest{r, err.Error()}, nil
		//}
		return accesstokenResponse{token, ""}, nil
	}
}

func DecodeAuthorizeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req authorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeAccessTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req accesstokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type authorizeRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type authorizeResponse struct {
	Data struct {
		Authorization_code string
		Expires_at         int
	}
	Status int
}

type accesstokenRequest struct {
}

type accesstokenResponse struct {
}
