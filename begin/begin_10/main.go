package main

import "fmt"

func main() {
	var a float32
	var b float32
	fmt.Println("A = ")
	fmt.Scan(&a)
	fmt.Println("B = ")
	fmt.Scan(&b)
	fmt.Println("Сумма чисел:", a*a+b*b)
	fmt.Println("Разность чисел:", a*a-b*b)
	fmt.Println("Произведение чисел:", a*a*b*b)
	fmt.Println("Частность чисел:", (a*a)/(b*b))
}
