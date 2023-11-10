package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите длину А: ")
	fmt.Scan(&a)
	fmt.Println("Введите длину B: ")
	fmt.Scan(&b)
	fmt.Println("В отрезке А находится ", a/b, "полных отрезков B")
}
