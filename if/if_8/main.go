package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println("Число ", a, " больше чем число ", b)
	} else if b < a {
		fmt.Println("Число ", b, " больше чем число ", a)
	} else {
		fmt.Println("Числа равны")
	}
}
