package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ncostamagna/axul_notifications/internal/notification"
	"github.com/ncostamagna/axul_notifications/pkg/handler"
)

func main() {

	fmt.Println("Entraaaaa")
	r := notification.NewRepository(nil, nil)
	s := notification.NewService(r, nil)
	e := notification.MakeEndpoints(s)
	h := handler.NewLambdaNotificationGetAll(e)
	lambda.StartHandler(h)
}
