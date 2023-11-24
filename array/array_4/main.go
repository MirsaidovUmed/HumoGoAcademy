package main

import "fmt"

func main() {
	var numArr int
	fmt.Scan(&numArr)
	arr := make([]int, numArr)
	count := 0
	for i := 0; i < numArr; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 1; i < numArr-1; i++ {
		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			count++
		}
	}
	fmt.Println(count)
}
