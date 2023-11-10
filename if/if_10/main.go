package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a != b {
		result := a + b
		a = result
		b = result
	} else {
		a = 0
		b = 0
	}
	fmt.Println("Новое значение А", a)
	fmt.Println("Новое значение B", b)
}
