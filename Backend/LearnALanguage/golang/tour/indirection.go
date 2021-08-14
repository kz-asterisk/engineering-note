package main

import (
	"fmt"
)

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Sacle(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Sacle(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Sacle(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
