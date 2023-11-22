package main

import "fmt"

func main() {
	var N int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	K := 1
	val := 3
	for val <= N {
		K++
		val = val * 3
	}

	fmt.Println("Наименьшее целое число K, при котором выполняется неравенство 3^K > N =", K)
}
