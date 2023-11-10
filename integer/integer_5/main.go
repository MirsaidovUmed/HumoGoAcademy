package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите длину А: ")
	fmt.Scan(&a)
	fmt.Println("Введите длину B: ")
	fmt.Scan(&b)
	fmt.Println("На отрезке A не занято: ", a%b)
}
