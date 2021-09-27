package main

import (
	"fmt"
	"ikimono/ikimono"
)

func init() {
	fmt.Println("init処理の実行")
}

func main() {
	okami := ikimono.NewIkimono("okami", 10, 10, 10)
	fmt.Println(okami.ShowIkimonoName())
	ezoshika := ikimono.NewIkimonoOnlyName("ezoshika")
	fmt.Println(ezoshika.ShowIkimonoName())
	fmt.Println(okami.Name, "の攻撃で", ezoshika.Name, "は", ikimono.CalcDamage(okami, ezoshika), "のダメージ❕")
}
