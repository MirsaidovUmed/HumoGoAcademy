package main

import "fmt"

func main() {
	var N int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	K := 0
	val := 1
	for val < N {
		K++
		val = val * 3
	}

	if val >= N {
		K--
	}

	fmt.Println("Наибольшее целое число K, при котором выполняется неравенство 3^K < N = ", K)
}
