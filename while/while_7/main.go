package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var K int
	for K = 1; K*K <= N; K++ {
	}

	fmt.Println("Наименьшее целое положительное число K, квадрат которого превосходит N = ", K)
}
