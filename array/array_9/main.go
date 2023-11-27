package main

import "fmt"

func main() {
	var sizeArr, numArr, check int
	fmt.Scan(&sizeArr)
	arr := make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&numArr)
	for i := 0; i < sizeArr; i++ {
		if arr[i] == numArr {
			check = i + 1
			fmt.Println(check)
		}
	}
}
