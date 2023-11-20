package main

import "fmt"

func main() {
	var n, result int
	fmt.Scan(&n)
	for i := n; i <= 2*n; i++ {
		result += i * i
	}
	fmt.Println(result)
}
