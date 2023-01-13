package mysql

import (
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

func QueryConfessionByID(confession_id string) (*model.Confession, error) {
	var confession model.Confession
	Client.Table(TableName).Where("confession_id = ?", confession_id).Scan(&confession)
	return &confession, nil
}

func QueryConfessionByUser(user_id string) (*model.Confession, error) {
	var confession model.Confession
	Client.Table(TableName).Where("user_id = ?", user_id).Scan(&confession)
	return &confession, nil
}

func Save(conffession *model.Confession) error {
	Client.Table(TableName).Save(conffession)
	return nil
}
