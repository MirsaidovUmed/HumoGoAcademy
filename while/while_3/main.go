package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k)
	for n >= k {
		n -= k
		m++
	}
	fmt.Println("Частное от деления нацело N на K =", m, "Остаток от этого деления =", n)
}
