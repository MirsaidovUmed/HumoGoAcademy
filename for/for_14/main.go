package main

import "fmt"

func main() {
	var (
		n   int
		res int = 0
	)
	fmt.Scan(&n)

	for i := 1; i <= (2*n + 1); i += 2 {
		res += i
	}

	fmt.Println(res)
}
