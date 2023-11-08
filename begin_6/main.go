package main

import "fmt"

func main() {
	var firstNum int
	var secondNum int
	var thirdNum int
	fmt.Println("А = ")
	fmt.Scan(&firstNum)
	fmt.Println("B = ")
	fmt.Scan(&secondNum)
	fmt.Println("C = ")
	fmt.Scan(&thirdNum)
	fmt.Println("V = ", firstNum*secondNum*thirdNum)
	fmt.Println("S = ", 2*(firstNum*secondNum+secondNum*thirdNum+firstNum*thirdNum))
}
