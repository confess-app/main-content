package handler

import (
	"encoding/json"
	"fmt"
	"main-content/model"
	"main-content/service/mysql"
	"net/http"
	"strconv"

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

func GetByTag(query map[string]string, user *model.User) (events.APIGatewayProxyResponse, error) {
	pageStr, ok := query["page"]
	if !ok {
		pageStr = "1"
	}
	tag, ok := query["tag"]
	if !ok {
		tag = mysql.HotTag
	}
	limitStr, ok := query["limit"]
	if !ok {
		limitStr = "30"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Println("wrong value query string, page: " + pageStr)
		return CreateResponse("wrong value query string, page: "+pageStr, http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Println("wrong value query string, limit: " + limitStr)
		return CreateResponse("wrong value query string, limit: "+limitStr, http.StatusBadRequest)
	}
	confessions := []model.Confession{}
	if tag == mysql.NewTag {
		confessions, err = mysql.QueryConfessionWithPaging(page, limit)
		if err != nil {
			fmt.Println(err.Error())
			return CreateResponse(err.Error(), http.StatusInternalServerError)
		}
	}
	if len(confessions) == 0 {
		fmt.Println("confession post had end")
		return CreateResponse("confession post had end", http.StatusNotFound)
	}
	respBody, err := json.Marshal(confessions)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse(string(respBody), http.StatusOK)
}
