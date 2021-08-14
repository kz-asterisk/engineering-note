package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	//ポイントにすることでmain関数で指定した変数vを操作するためにはポインタにする必要あり
	fmt.Println(v.X, &v.X, v.Y, &v.Y, ":ポインタレシーバで定義したメソッドの中のv")
}

func (v Vertex) ScaleNonpointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, &v.X, v.Y, &v.Y, ":通常レシーバで定義したメソッドの中のv")
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.X, &v.X, v.Y, &v.Y, ":mainで定義したv")
	v.Scale(10)
	fmt.Println(v.Abs())
	v.ScaleNonpointer(4)
	fmt.Println(v.Abs())
}
