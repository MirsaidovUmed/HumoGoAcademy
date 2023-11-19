package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := 0
	for i := a; i <= b; i++ {
		sum += i * i
	}
	fmt.Println(sum)
}
