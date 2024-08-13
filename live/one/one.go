package one

import (
	"fmt"
	"go_study/live/two"
)

var B = NewMysqlClient("localhost.com")

type MysqlClient struct {
	Url string
}

func NewMysqlClient(url string) *MysqlClient {
	fmt.Println(two.C)
	return &MysqlClient{url}
}
