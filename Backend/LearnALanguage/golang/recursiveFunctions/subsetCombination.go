package main

import "fmt"

func subsetCombination(i int, w int, a []int) bool {
	fmt.Println("1:Case:", "i=", i, ",", "w=", w)
	if i == 0 {
		fmt.Println("検証")
		if w == 0 {
			fmt.Println("hit")
			return true
		} else {
			fmt.Println("hitせず")
			return false
		}
	}

	if subsetCombination(i-1, w, a) {
		fmt.Println("2:Case:", "i=", i, ",", "w=", w)
		return true
	}

	if subsetCombination(i-1, w-a[i-1], a) {
		fmt.Println("3:Case:", "i=", i, ",", "w=", w)
		return true
	}

	return false
}

func main() {
	var N, W int
	fmt.Scan(&N, &W)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	if subsetCombination(N, W, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
