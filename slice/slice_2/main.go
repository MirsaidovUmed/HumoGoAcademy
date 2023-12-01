package main

import "fmt"

func main() {
	var sizeOfArr int
	fmt.Scan(&sizeOfArr)
	arr := make([]int, sizeOfArr)

	for i := 0; i < sizeOfArr; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i+1], arr[sizeOfArr-1] = arr[sizeOfArr-1], arr[i+1]
	}
	fmt.Println(arr)
}
