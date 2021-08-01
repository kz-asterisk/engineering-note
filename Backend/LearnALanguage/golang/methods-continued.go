package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type ChomeString string

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (c ChomeString) SayHOO() string {
	msg := string(c) + " " + "HOO!!🎸"
	return msg
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
	c := ChomeString("Say")
	fmt.Println(c.SayHOO())

}
