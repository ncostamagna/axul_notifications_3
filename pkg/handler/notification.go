package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalhouse-dev/dh-kit/response"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/awslambda"
	"github.com/ncostamagna/axul_notifications/internal/notification"
	"gorm.io/gorm"
)

func NewLambdaNotificationGetAll(endpoints notification.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.GetAll), decodeProdUserStoreRequestHandler, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func decodeProdUserStoreRequestHandler(_ context.Context, payload []byte) (interface{}, error) {
	var gateway events.APIGatewayProxyRequest
	err := json.Unmarshal(payload, &gateway)
	if err != nil {
		return nil, response.InternalServerError("")
	}
	var event events.SNSEvent
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return nil, response.InternalServerError("")
	}

	var body string
	switch {
	case gateway.Body != "":
		body = gateway.Body

	case len(event.Records) > 0 && event.Records[0].SNS.Message != "":
		body = event.Records[0].SNS.Message
	default:
		return nil, response.BadRequest("")
	}

	var res notification.GetAllRequest
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func EncodeResponse(_ context.Context, resp interface{}) ([]byte, error) {
	var res response.Response
	switch resp.(type) {
	case response.Response:
		res = resp.(response.Response)
	default:
		res = response.InternalServerError("unknown response type")
	}
	return APIGatewayProxyResponse(res)
}

func APIGatewayProxyResponse(res response.Response) ([]byte, error) {
	bytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	awsResponse := events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: res.StatusCode(),
		Headers:    res.GetHeaders(),
	}
	return json.Marshal(awsResponse)
}

func HandlerErrorEncoder(log log.Logger) awslambda.HandlerOption {
	return awslambda.HandlerErrorEncoder(
		awslambda.ErrorEncoder(errorEncoder(log)),
	)
}

func HandlerFinalizer(log log.Logger) func(context.Context, []byte, error) {
	return func(ctx context.Context, resp []byte, err error) {
		if err != nil {
			log.Log("err", err)
		}
	}
}

func errorEncoder(log log.Logger) func(context.Context, error) ([]byte, error) {
	return func(_ context.Context, err error) ([]byte, error) {
		res := buildResponse(err, log)
		return APIGatewayProxyResponse(res)
	}
}

func buildResponse(err error, log log.Logger) response.Response {
	switch err.(type) {
	case response.Response:
		return err.(response.Response)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("")
	}
	log.Log("err", err)
	return response.InternalServerError("")
}
