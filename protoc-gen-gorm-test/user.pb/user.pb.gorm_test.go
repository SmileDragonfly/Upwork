package user_pb

import (
	"context"
	"github.com/jinzhu/gorm"
	"testing"
)

func ConnectDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=123@123A dbname=gorm port=5432 sslmode=disable"
	conn, err := gorm.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return conn
}

func TestUserServiceDefaultServer_Create(t *testing.T) {
	conn := ConnectDB(t)
	server := UserServiceDefaultServer{DB: conn}
	t.Run("Create", func(t *testing.T) {
		resp, err := server.Create(context.Background(), &CreateUserRequest{Payload: &User{
			Id:          1,
			Name:        "Dat",
			Emails:      "dat.doantrongk57@gmail.com",
			HomeAddress: "Hoang Mai, Hanoi",
			PhoneNumber: "0123456789",
			Description: "freelance",
		}})
		if err != nil {
			t.Fatal(err)
			return
		}
		print(resp.String())
	})
}

func TestUserServiceDefaultServer_Read(t *testing.T) {
	conn := ConnectDB(t)
	server := UserServiceDefaultServer{DB: conn}
	t.Run("Read", func(t *testing.T) {
		resp, err := server.Read(context.Background(), &ReadUserRequest{Id: 1})
		if err != nil {
			t.Fatal(err)
			return
		}
		print(resp.String())
	})
}

func TestUserServiceDefaultServer_Update(t *testing.T) {
	conn := ConnectDB(t)
	server := UserServiceDefaultServer{DB: conn}
	t.Run("Update", func(t *testing.T) {
		resp, err := server.UpdateUser(context.Background(), &UpdateUserRequest{
			Payload: &User{
				Id:          1,
				Name:        "1212",
				Emails:      "321312",
				HomeAddress: "31231",
				PhoneNumber: "12313",
				Description: "2313",
			},
		})
		if err != nil {
			t.Fatal(err)
			return
		}
		print(resp.String())
	})
}

func TestUserServiceDefaultServer_Delete(t *testing.T) {
	conn := ConnectDB(t)
	server := UserServiceDefaultServer{DB: conn}
	t.Run("Delete", func(t *testing.T) {
		resp, err := server.Delete(context.Background(), &DeleteUserRequest{Id: 1})
		if err != nil {
			t.Fatal(err)
			return
		}
		print(resp.String())
	})
}
