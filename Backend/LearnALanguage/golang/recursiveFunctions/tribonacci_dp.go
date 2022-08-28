package main

import "fmt"

var memo []int

func tribonacciDp(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 0
	} else if N == 2 {
		return 1
	}

	if memo[N] != -1 {
		return memo[N]
	}

	memo[N] = tribonacciDp(N-1) + tribonacciDp(N-2) + tribonacciDp(N-3)
	return memo[N]
}

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		memo = append(memo, -1)
	}

	fmt.Println(tribonacciDp(N - 1))
}
