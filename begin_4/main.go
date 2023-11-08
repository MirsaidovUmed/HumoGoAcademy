package main

import "fmt"

func main() {
	var d float32
	var p float32 = 3.14

	fmt.Println("Введите диаметр:")

	fmt.Scan(&d)

	fmt.Println("Длинна окружности = ", p*d)
}
