package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println(num + 1)
	} else {
		fmt.Println("Ладно можно тронут сделаем минус два num =", num-2)
	}
}
