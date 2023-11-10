package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите двузначное число:")
	fmt.Scan(&num)
	fmt.Println("Перестановка чисел = ", (num/10)+(num%10)*10)
}
