package user

import (
	"errors"
	"fmt"
	"main-content/internal/handler"
	"main-content/model"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func AuthenUser(req events.APIGatewayProxyRequest) (*model.User, error) {
	parseHeader := http.Request{Header: http.Header{"Cookie": []string{req.Headers["Cookie"]}}}
	tokenCookie, err := parseHeader.Cookie("token")
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("unauthorized")
	}
	user, err := handler.DecodeTokenToUserModel(tokenCookie.Value)
	fmt.Printf("user: %+v\n", user)
	if err != nil {
		return nil, errors.New("token had expired")
	}
	return user, nil
}
