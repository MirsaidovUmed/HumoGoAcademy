package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println(num + 1)
	} else if num < 0 {
		fmt.Println("Ладно можно тронут сделаем минус два num =", num-2)
	} else {
		fmt.Println("Теперь ещё и меняем 0 на", num+10)
	}
}
