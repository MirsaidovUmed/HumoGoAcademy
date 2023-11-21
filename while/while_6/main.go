package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 1.0
	for n > 0 {
		result *= float64(n)
		n -= 2
	}
	fmt.Println(result)
}
