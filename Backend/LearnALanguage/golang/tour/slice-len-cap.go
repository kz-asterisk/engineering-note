package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[:2]
	printSlice(s)

	//スライスの容量を超えて伸ばそうとするとエラーになる
	//panic: runtime error: slice bounds out of range [:10] with capacity 6
	//s = s[:10]
	//printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
