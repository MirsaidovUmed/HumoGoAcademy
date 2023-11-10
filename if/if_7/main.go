package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println("Порядковый номер числа:1")
	} else if b < a {
		fmt.Println("Порядковый номер числа:2")
	} else {
		fmt.Println("Числа равны")
	}
}
