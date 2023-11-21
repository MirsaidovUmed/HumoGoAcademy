package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	for n%2 == 0 {
		n = n / 2
		k++
	}
	fmt.Println(k)
}
