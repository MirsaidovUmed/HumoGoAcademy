package main

import "fmt"

func main() {
	var a, b, c int
	var d = 0
	var e = 0
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		d = d + 1
	} else if a < 0 {
		e = e + 1
	}
	if b > 0 {
		d = d + 1
	} else if b < 0 {
		e = e + 1
	}
	if c > 0 {
		d = d + 1
	} else if c < 0 {
		e = e + 1
	}
	fmt.Println("Кол-во чётных чисел", d, "Кол-во нечётных чисел", e)
}
