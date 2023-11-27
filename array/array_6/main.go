package main

import "fmt"

func main() {
	var sizeOfArr, numInArray, count int
	fmt.Scan(&sizeOfArr)
	arr := make([]int, sizeOfArr)
	for i := 0; i < sizeOfArr; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&numInArray)
	for i := 0; i < sizeOfArr; i++ {
		if arr[i] == numInArray {
			count++
		}
	}
	fmt.Println(count)

}
