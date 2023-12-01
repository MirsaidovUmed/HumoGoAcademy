package main

import (
	"fmt"
)

func main() {
	var (
		n int
		k int
	)

	fmt.Scan(&n, &k)
	numbers := make([]int, n)

	for i := 0; i < k; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < n; i += k {
		end := i + k
		if end > k {
			end = k
		}
		fmt.Println(numbers[i:end])
	}
}
