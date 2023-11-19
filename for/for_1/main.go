package main

import "fmt"

func main() {
	//даны числа K N
	//K - 3
	//N - 0, 1, 2
	var N, K int
	fmt.Scan(&N, &K)
	for i := 0; i < N; i++ {
		fmt.Println(K)
	}
}
