package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoimpl"
	models "go-postgres/model"
)

const (
	host = "localhost"
	port = "5432"
	user = "huyctdev"
	password = "password"
	dbname = "todo"
)

func main(){
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	userRepo := repoimpl.NewUserRepo(db.SQL)

	user1 := models.User{
		ID: 1,
		Name: "huyct",
		Gender: "Male",
		Email: "huy@gmail.com",
	}

	user2 := models.User{
		ID: 2,
		Name: "huyct1",
		Gender: "Male",
		Email: "huy1@gmail.com",
	}

	userRepo.Insert(user1)
	userRepo.Insert(user2)

	users,_ := userRepo.Select()

	for i:= range users {
		fmt.Println(users[i])
	}
}