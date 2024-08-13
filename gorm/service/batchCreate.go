package service

import (
	"go_study/gorm/global"
	"go_study/gorm/model"
	"log"
	"strconv"
)

type Service struct{}

func (Service) BatchCreate() {
	var tags []model.Tag
	for i := 8; i < 13; i++ {
		var t model.Tag
		t.ID = uint(i)
		t.TagName = "tag6666 " + strconv.Itoa(i)
		tags = append(tags, t)

	}

	var users []model.User
	for i := 0; i < 4; i++ {
		var p model.User
		p.ID = uint(i + 1)
		p.Age = i
		p.Username = "ximi " + strconv.Itoa(i)
		p.Tags = tags
		users = append(users, p)
	}

	err := global.DB.Debug().Save(&users)
	if err != nil {
		log.Fatal(err.Error)
	}
	log.Println("okkk")

}
