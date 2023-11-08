package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Длина куба = ")
	fmt.Scan(&num)
	fmt.Println("V = ", num*num*num)
	fmt.Println("S = ", 6*num*num)
}
