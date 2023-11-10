package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("A = ")
	fmt.Scan(&a)
	fmt.Println("B = ")
	fmt.Scan(&b)
	fmt.Println("Среднее геометрическое А и В: ", math.Sqrt(a*b))
}
