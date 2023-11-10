package main

import "fmt"

func main() {
	var firstNum int

	var secondNum int

	fmt.Println("Сторона прямоугольника А = ")

	fmt.Scan(&firstNum)

	fmt.Println("Сторона прямоугольника B = ")

	fmt.Scan(&secondNum)

	fmt.Println("Периметр прямоугольника = ", 2*(firstNum+secondNum))

	fmt.Println("Площадь прямоугольника = ", firstNum*secondNum)
}
