package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n       int
		results []int
		minPos  = math.MaxInt
		maxNeg  = math.MinInt
	)
	fmt.Scan(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < n; i++ {
		if numbers[i] < 0 {
			if maxNeg < numbers[i] {
				maxNeg = numbers[i]
			}
		}

		if numbers[i] > 0 {
			if minPos > numbers[i] {
				minPos = numbers[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if numbers[i] != maxNeg && numbers[i] != minPos {
			results = append(results, numbers[i])
		}
	}

	fmt.Println(results)
}
