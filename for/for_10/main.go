package main

import "fmt"

func main() {
	var (
		n      int
		result float32
	)

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		result += 1 / float32(i)
	}
	fmt.Println(result)
}
