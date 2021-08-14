package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	//ch <- 3
	//fmt.Println(<-ch) -> 途中で取り出すとbufferが溢れない
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
