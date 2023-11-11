package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println((a < b && b < c) || (a > b && b > c))
}
