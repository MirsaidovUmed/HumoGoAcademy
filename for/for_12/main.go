package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	result := 1.0
	for i := 1.1; i <= n; i += 0.1 {
		result *= i
	}
	fmt.Println(result)
}
