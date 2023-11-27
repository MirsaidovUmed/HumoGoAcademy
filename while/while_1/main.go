package main

import "fmt"

func main() {
	var a, b int     // задаём две переменные
	fmt.Scan(&a, &b) // задаём две переменные
	for a >= b {     // пока a >= b
		a -= b //a будет отнимать от себя куски b
	}
	fmt.Println(a) //вывод оставшихся кусков
}
