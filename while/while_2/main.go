package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for a >= b {
		a -= b
		c++
	}
	fmt.Println("Количество отрезков B =", c)
}
