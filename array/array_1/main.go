package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 1; i < n; i++ {
		if (arr[i] > 0 && arr[i+1] > 0) || (arr[i+1] < 0 && arr[i] < 0) {
			fmt.Println("YES")
			break
		} else {
			fmt.Println("NO")
			break
		}
	}

}
