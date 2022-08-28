package main

import (
	"fmt"	
)

var memo []int

func fibo(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}

	if memo[N] != -1 {
		return memo[N]
	}

	memo[N] = fibo(N-1) + fibo(N-2)

	return memo[N]
}

func main() {
	for i := 0; i < 50; i++ {
		memo = append(memo, -1)
	}
	fibo(49)

	for N := 2; N < 50; N++ {
		fmt.Println(N, "項目:", memo[N])
	}
}
