package main

import (
	"fmt"
)

func main() {
	var price float64
	fmt.Scan(&price)
	for i := 1.2; i <= 2; i += 0.2 {
		fmt.Println(price * i)
	}
}
