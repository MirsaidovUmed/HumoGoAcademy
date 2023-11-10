package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите двузначное число:")
	fmt.Scan(&num)
	fmt.Println("Сложение левого и правого чисел = ", (num/10)+(num%10))
	fmt.Println("Произведение левого и правого чисел = ", (num/10)*(num%10))
}
