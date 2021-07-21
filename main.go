package main

import "fmt"

func sum(a, b, c int) int {
	return a + b + c
}
func sub(a, b, c int) int {
	return a - b - c
}

func main() {
	fmt.Println("Sum: ", sum(1, 2, 3))
	fmt.Println("Sub: ", sub(10, 2, 3))
}
