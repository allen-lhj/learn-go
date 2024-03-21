package main

import (
	"log"

	"github.com/allen-lhj/myprogram/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// PrintText("Hi")
	// 创建一个传递int类型的channel，可以用来发送和接收整形数据
	intChan := make(chan int)
	// 确保在main函数结束前关闭channel，防止内存泄露
	defer close(intChan)
	// 启动一个新的goroutine来执行Cal函数
	go CalculateValue(intChan)
	// 通过num进行接收
	num := <-intChan
	log.Println(num)
}

// func PrintText(s string) {
// 	log.Println(s)
// }
