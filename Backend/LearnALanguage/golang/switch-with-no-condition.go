package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("type: %T value: %v\n", t, t)
	fmt.Printf("type: %T value: %v\n", t.Hour(), t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
