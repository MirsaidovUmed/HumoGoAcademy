package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	if a > b {
		c = a
		a = b
		b = c
		fmt.Println(a, b)
	} else if a < b {
		//вроде как этой части в условии не было, но добавлю что это при условии что
		//B  больше чем А они так же должны меняться
		c = a
		a = b
		b = c
		fmt.Println(a, b)
	}
}
