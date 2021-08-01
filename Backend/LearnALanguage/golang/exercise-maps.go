package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	for _, word := range strings.Fields(s){
		//fmt.Println(word)
		result[word]++
		//fmt.Println(result)
	}
	return result
}

func main() {
	result := WordCount("I am a Pen I Love Music")
	fmt.Println(result)
}
