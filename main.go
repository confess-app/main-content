package main

import (
	"main-content/service/mysql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	PostPath = "/content/post"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod == http.MethodGet {
		return events.APIGatewayProxyResponse{
			Body:       "error: method get not allowed with this url",
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}

	switch req.Path {
	case PostPath:
		mysql.Init()
		return events.APIGatewayProxyResponse{
			Body:       "error: url/api not exists",
			StatusCode: http.StatusNotFound,
		}, nil
		// return handler.Register(req.Body)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "error: url/api not exists",
			StatusCode: http.StatusNotFound,
		}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest(events.APIGatewayProxyRequest{
	// 	HTTPMethod: http.MethodPost,
	// 	Path:       LoginPath,
	// 	Body:       "{ \"username\": \"duong7\", \"password\": \"password_duong\" }",
	// })
}
