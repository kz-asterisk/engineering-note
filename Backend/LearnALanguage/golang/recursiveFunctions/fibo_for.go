package main

import "fmt"

func main(){
	F := make([]int,50)
	F[0] = 0
	F[1] = 1
	for N := 2; N < 50; N++{
		F[N] = F[N - 1] + F[N - 2]
		fmt.Println(N,"項目:",F[N])
	}
}