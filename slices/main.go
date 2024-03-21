package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string
	var mySlice2 []int
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Bob")
	mySlice = append(mySlice, "Alice")

	mySlice2 = append(mySlice2, 3)
	mySlice2 = append(mySlice2, 4)
	mySlice2 = append(mySlice2, 1)

	log.Println("mySlice is set to", mySlice)
	log.Println("mySlice2 is set to", mySlice2)

	// 排序
	sort.Ints(mySlice2)
	numbers := []int{1, 2, 3, 4, 5}

	// 截取
	log.Println((numbers[:3]))

}
