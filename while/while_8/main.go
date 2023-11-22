package main

import "fmt"

func main() {
	var N int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	var K int
	for K = 1; K*K <= N; K++ {
	}

	K = K - 1

	fmt.Println("Наибольшее целое число K, квадрат которого не превосходит N =", K)
}
