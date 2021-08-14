package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string){
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	words := []string{"Hello", "Bye", "HOO", "Yeah!"}
	fmt.Println("Begin")
	for _, w := range words{
		wg.Add(1)
		go say(w)
	}
	wg.Wait()
	fmt.Println("End")
}
