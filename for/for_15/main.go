package main

import "fmt"

func main() {
	var (
		n   int
		a   float64
		res float64
	)
	fmt.Scan(&a, &n)

	for i := 1; i <= n; i++ {
		res *= float64(i)
	}

	fmt.Println(res)

}
