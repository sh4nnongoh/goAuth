package registration

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sh4nnongoh/goAuth/src/user"
)

func MakeRegistrationEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(registrationRequest)
		svc.New(req.ID, req.Email, req.Name)
		return registrationResponse{Status: 1}, nil
	}
}

type registrationRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Name     user.Name
	Email    user.Email
	ID       user.UserID
}

type registrationResponse struct {
	Status int
}
