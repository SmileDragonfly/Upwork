package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	user_pb "gormgen1/user.pb"
)

func main() {
	dsn := "host=localhost user=postgres password=123@123A dbname=gorm port=5432 sslmode=disable"
	conn, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.AutoMigrate(&user_pb.User{})
	_ = user_pb.UserServiceDefaultServer{DB: conn}
}
