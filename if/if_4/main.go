package main

import "fmt"

func main() {
	var a, b, c int
	var d = 0
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		d = d + 1
	}
	if b > 0 {
		d = d + 1
	}
	if c > 0 {
		d = d + 1
	}
	fmt.Println(d)
}
