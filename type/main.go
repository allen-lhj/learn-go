package main

import (
	"log"
	"time"
)

type User struct {
	FristName string
	LastName  string
	Phone     string
	Age       int
	Birthday  time.Time
}

func main() {
	user := User{
		FristName: "John",
		LastName:  "Doe",
		Phone:     "123-456-7890",
		Birthday:  time.Date(1980, 1, 2, 0, 0, 0, 0, time.UTC),
	}
	log.Println(user)
}
