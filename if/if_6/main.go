package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("Число а больше чем б", a)
	} else if a < b {
		fmt.Println("Число б больше чем а", b)
	}
}
