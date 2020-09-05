package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

// creamos function para crear Endpoinst struct
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// casteamos a CreateUserRequest
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)

		// respondemos en base al response que definimos
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email: email,
		}, err
	}
}
