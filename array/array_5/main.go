package main

import "fmt"

func main() {
	var n, min_ int

	fmt.Scan(&n)

	array_ := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array_[i])

		if array_[i]%2 != 0 && (min_ == 0 || array_[i] < min_) {
			min_ = array_[i]
		}
	}

	fmt.Println(min_)
}
