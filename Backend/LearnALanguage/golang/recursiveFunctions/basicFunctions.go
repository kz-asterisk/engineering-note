package main

import (
	"fmt"
)

func recursive(N int) int {
	fmt.Println(N,"を呼び出しました")	

	if N == 0 {
		return 0
	}

	result := N + recursive(N -1)
	fmt.Println(N,"までの和 = ",result)

	return result
}

func main() {
	recursive(5)
}