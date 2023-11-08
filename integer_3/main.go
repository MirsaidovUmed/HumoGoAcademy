package main

import "fmt"

func main() {
	var byte int
	fmt.Println("Byte = ")
	fmt.Scan(&byte)
	fmt.Println(byte/1024, "Килобайт")
}
