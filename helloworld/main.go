package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	var whatToSay string
	var i int
	whatToSay = "Goodbye, cruel world!"
	i = 7
	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)
	whatWasSaid := saySomething()
	fmt.Println("The function said:", whatWasSaid)
}

func saySomething() string {
	return "something"
}
