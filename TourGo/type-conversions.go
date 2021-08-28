package main

import (
	"fmt"
	"math"
)

func simple() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(u)
}

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}

//var 変数名　型名　or 変数名 := はしっかりする　＝＞　静的型付け言語だからちゃんと型宣言した方が良さそう
