package main

import "fmt"

func fibo(N int) int {
	var result int
	fmt.Printf("fibo(%v)を呼び出しました\n",N)
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	} else {
		result = fibo(N-1) + fibo(N-2)
		fmt.Printf("%v項目 = %v\n",N,result)
	}

	return result
}

func main() {
	fibo(6)
}
