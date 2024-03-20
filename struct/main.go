package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar1 myStruct
	myVar1.FirstName = "John"

	var myVar2 myStruct
	myVar2.FirstName = "Bob"

	log.Println("myVar is set to", myVar1.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
