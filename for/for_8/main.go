package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := 1
	for i := a; i <= b; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
