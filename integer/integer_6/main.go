package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите двузначное число:")
	fmt.Scan(&num)
	fmt.Println("Левая цифра", num/10)
	fmt.Println("Правая цифра", num%10)
}
