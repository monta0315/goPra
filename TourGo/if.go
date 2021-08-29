package main

import (
	"fmt"
	"math"
)

func sqrt(x float64)string  {
	if x<0{
		return sqrt(-x) + "i"
	}
	return string(math.Sqrt(x))
}

func main()  {
	fmt.Println(sqrt(2),sqrt(-4))
}
//Sprintで文字列をリターンしている=>型が分からなくても使えるそう
