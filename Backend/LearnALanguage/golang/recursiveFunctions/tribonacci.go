package main

import "fmt"

func tribonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 0
	} else if i == 2 {
		return 1
	}

	return tribonacci(i-1) + tribonacci(i-2) + tribonacci(i-3)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(tribonacci(N-1))
}
