package main

import (
	"fmt"
)

func Sqrt(x float64)float64{
	z := 1.0
	time := 0
	for time<11{
		z -= (z*z-x)/2*z
		time++
	}
	return z
}

func main()  {
	fmt.Println(Sqrt(2))
}
