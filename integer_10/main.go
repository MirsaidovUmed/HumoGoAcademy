package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите трёхзначное число")
	fmt.Scan(&num)
	fmt.Println("Последняя цифра = ", num%10)
	fmt.Println("Средняя цифра = ", num%100/10)
}
