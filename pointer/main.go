package main

import "log"

func main() {
	var myString string

	myString = "Red"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Blue"
	*s = newValue
}
