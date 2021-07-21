package main

import "fmt"

func sum(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("Sum: ", sum(1, 2, 3))
}
