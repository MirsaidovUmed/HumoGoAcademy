package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 1
	for i := 1; i <= n; i++ {
		result = 1 / n
	}
	fmt.Println(float64(result))

}
