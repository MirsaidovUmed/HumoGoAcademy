package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите трёхзначное число")
	fmt.Scan(&num)
	fmt.Println("Первая цифра = ", num/100)
}
