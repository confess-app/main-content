package handler

import (
	"fmt"
	"main-content/model"
	"main-content/service/mysql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

func PostContent(body string, user *model.User) (events.APIGatewayProxyResponse, error) {
	data := ConfessionData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	newConfID, err := uuid.NewRandom()
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	confession := &model.Confession{
		ConfessionID: newConfID.String(),
		UserID:       user.UserID,
		Username:     user.Username,
		Anonymous:    data.Anonymous,
		Content:      data.Content,
		Category:     data.Category,
	}
	err = mysql.Save(confession)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse(newConfID.String(), http.StatusOK)
}

type ConfessionData struct {
	Anonymous bool   `json:"anonymous"`
	Content   string `json:"content"`
	Category  string `json:"category"`
}
