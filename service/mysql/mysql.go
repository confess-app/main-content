package mysql

import (
	"fmt"
	"main-content/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TableName = "confession"
)

var Client *gorm.DB

func Init() {
	conn := os.Getenv("MYSQL_DATABASE_URI")
	var err error
	Client, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = Client.AutoMigrate(&model.Confession{})

	if err != nil {
		panic(err)
	}
}

func QueryConfessionByID(confessionID string) (*model.Confession, error) {
	var confession model.Confession
	Client.Table(TableName).Where("confession_id = ?", confessionID).Scan(&confession)
	return &confession, nil
}

func QueryConfessionByUser(userID string) (*model.Confession, error) {
	var confession model.Confession
	Client.Table(TableName).Where("user_id = ?", userID).Scan(&confession)
	return &confession, nil
}

func Save(conffession *model.Confession) error {
	Client.Table(TableName).Save(conffession)
	return nil
}

func Close() {
	db, err := Client.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.Close()
	if err != nil {
		fmt.Println("cannot close connection: ", err.Error())
	}
	fmt.Println("connection closed")
}
