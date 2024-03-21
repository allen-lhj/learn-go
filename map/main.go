package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

// map 的存储总是随机的，你不能按照放进去的顺序去查找他们
func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "John",
		LastName:  "Sawler",
	}
	myMap["me"] = me
	log.Println(myMap["me"])
}
