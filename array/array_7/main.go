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
	for i := 0; i <= sizeOfArr-1; i++ {
		if arr[i] == numInArray {
			count++
		}
	}
	if count > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
