package main

import "fmt"

func main() {
	A := []int{1, 4, 5, 8}
	B := []int{4, 5, 9, 10}
	C := []int{4, 6, 7, 10, 2}
	fmt.Println(Solution(A, B, C)) // expect 4
}
