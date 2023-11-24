package main

import "fmt"

func main() {
	var (
		n      int
		minArr = 0
	)
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	minArr = arr[n-1]

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			if minArr > arr[i] {
				minArr = arr[i]
			}
		} else {
			minArr = 0
		}
	}
	fmt.Println(minArr)
}
