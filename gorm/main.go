package main

import (
	"go_study/gorm/global"
	"go_study/gorm/service"
)

func main() {
	global.InitDB()
	//err := global.DB.AutoMigrate(
	//	model.User{},
	//	model.Tag{},
	//	model.Book{},
	//)
	//if err != nil {
	//	panic(err)
	//}
	ser := service.Service{}
	ser.BatchCreate()
}
