package main

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/model"
	"book-mgr-backend/routers"
)

func init() {
	dao.InitMysqlServer()

	if err := dao.Db.Model(&model.Book{}).AutoMigrate(&model.Book{}); err != nil {
		panic(err)
	}
	if err := dao.Db.Model(&model.User{}).AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	if err := dao.Db.Model(&model.History{}).AutoMigrate(&model.History{}); err != nil {
		panic(err)
	}

}

func main() {
	var app routers.App
	app.RunServer()
}
