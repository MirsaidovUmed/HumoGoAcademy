package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a)
	for fmt.Scan(&b); a <= b; a++ {
		fmt.Println(a)
		n++
	}
	fmt.Println("Количество n чисел:", n)
}
