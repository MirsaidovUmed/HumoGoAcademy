package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num/100 > num%10 {
		fmt.Println("Первое больше", num/100)
	} else {
		fmt.Println("Третье больше", num%10)
	}
}
