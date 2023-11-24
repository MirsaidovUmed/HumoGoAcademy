package main

import "fmt"

func main() {
	var (
		numArr     int
		biggestNum int
	)
	fmt.Scan(&numArr)
	arr := make([]int, numArr)
	for i := 0; i < numArr; i++ {
		fmt.Scan(&arr[i])
		if arr[i]%2 != 0 {
			if i == 0 || arr[i] < biggestNum {
				biggestNum = arr[i]
			}
		} else {
			fmt.Println(0)
			break
		}
	}
	fmt.Println(biggestNum)
}

//var n, biggestNum int
//fmt.Scan(&n)
//arr := make([]int, n)
//for i := 0; i < n; i++ {
//	fmt.Scan(&arr[i])
//}
//for i := 0; i < n-1; i++ {
//	if arr[i] < arr[i+1] {
//		biggestNum = arr[i+1]
//	} else {
//		biggestNum = arr[i]
//		break
//	}
//}
//fmt.Println(biggestNum)
