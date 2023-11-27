package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	marks := make([]int, n)
	sortedMarks := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&marks[i])
		sortedMarks[i] = marks[i]
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortedMarks[j] > sortedMarks[j+1] {
				sortedMarks[j], sortedMarks[j+1] = sortedMarks[j+1], sortedMarks[j]
			}
		}
	}

	min_ := sortedMarks[0]
	max_ := sortedMarks[n-1]

	for i := 0; i < n; i++ {
		if marks[i] == max_ {
			marks[i] = min_
		}
	}

	fmt.Println(marks)
}
