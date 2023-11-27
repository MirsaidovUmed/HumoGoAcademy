package main

import "fmt"

func main() {
	var sizeArr, numArr int
	fmt.Scan(&sizeArr)
	arr := make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&numArr)

	closest := arr[0]
	minDiff := closest - numArr
	if minDiff < 0 {
		minDiff = -minDiff
	}

	for i := 1; i < sizeArr; i++ {
		diff := arr[i] - numArr
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
			closest = arr[i]
		}
	}

	fmt.Println(closest)
}
