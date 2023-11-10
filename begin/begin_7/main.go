package main

import "fmt"

func main() {
	var r float32
	var p float32 = 3.14
	fmt.Println("R = ")
	fmt.Scan(&r)
	fmt.Println("L = ", 2*p*r)
	fmt.Println("S = ", p*r*r)
}
