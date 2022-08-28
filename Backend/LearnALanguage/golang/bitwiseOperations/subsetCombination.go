//集合(要素数N)から部分集合を作った際その合計が任意の数(W)になるかを算出する(線形探索)
package main

import (
	"fmt"
)

func main() {
	var N, W int
	fmt.Scan(&N, &W)

	A := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	EXIST := false

	//1 << N を設定することで必要な要素数+1の数を算出できる(シフト演算しているけ)
	//2進数の各桁を要素が「ある/ない」の選択に見立て、その組み合わせの総数を算出しているイメージ
	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			//現状の組み合わせとなるbitと各桁の1の論理積を計算(C++の場合は真偽値だがgolangの場合は論理積となる)
			if bit&(1<<i) != 0 {
				sum += A[i]
			}
		}
		if sum == W {
			EXIST = true
		}
	}
	fmt.Println(EXIST)
}
