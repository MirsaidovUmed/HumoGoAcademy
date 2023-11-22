package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("n=")
	fmt.Scan(&n)

	result := 0.0
	for i := 1; i <= n; i++ {
		result += math.Pow(-1, float64(i+1)) * (1.0 + float64(i)/10)
	}
	fmt.Printf("Result=%f", result)
}
