package main

import "fmt"

func main() {
	var num int

	fmt.Println("Введите сторону квадрата:")

	fmt.Scan(&num)

	fmt.Println("Площадь равна ", num*num)
}
