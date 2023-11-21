package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n%3 == 0 {
		fmt.Println(true)
		break
	}
	for n%3 != 0 {
		fmt.Println(false)
		break
	}

}
