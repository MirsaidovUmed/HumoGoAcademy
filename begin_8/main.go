package main

import "fmt"

func main() {
	var a float32
	var b float32
	fmt.Println("A = ")
	fmt.Scan(&a)
	fmt.Println("B = ")
	fmt.Scan(&b)
	fmt.Println("Среднее арифметическое А и В: ", (a+b)/2)
}
