package notification

import (
	"context"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	GetAllRequest struct {
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(s),
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Entra")
		fmt.Println(request)
		return response.OK("success", request, nil, nil), nil
	}
}
