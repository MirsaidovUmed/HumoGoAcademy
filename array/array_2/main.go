package main

import "fmt"

func main() {
	var (
		n    int // соклько число нужен
		temp int // для ввода элементы массива
	)
	fmt.Scan(&n) // input num

	mas := make([]int, n)     // созадем массив с длиной n
	for i := 1; i <= n; i++ { // ну обичный фор
		fmt.Scan(&temp) // число элементов массива
		mas[n-i] = temp // добавим число в массив начиная с конца до начало
	}

	fmt.Println(mas) // результат)
}
