package handler

import (
	"encoding/json"
	"fmt"
	"main-content/model"
	"main-content/service/mysql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetOne(query map[string]string, user *model.User) (events.APIGatewayProxyResponse, error) {
	confID, ok := query["confession_id"]
	if !ok {
		fmt.Println("wrong parameter, must confession_id")
		return CreateResponse("wrong parameter, must confession_id", http.StatusBadRequest)
	}
	conf, err := mysql.QueryConfessionByID(confID)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	if conf.ID == 0 {
		fmt.Println("confession post not found with id: " + confID)
		return CreateResponse("confession post not found with id: "+confID, http.StatusNotFound)
	}
	respBody, err := json.Marshal(conf)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse(string(respBody), http.StatusOK)
}

type GetOneData struct {
	ConfessionID string `json:"confession_id"`
}
