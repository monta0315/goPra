package main

import "fmt"

func main() {
	var i int
	j := i //j is an int
	k := 42
	f := 3.142
	g := 0.867 + 0.5i
	fmt.Printf("%T %v %T %v %T %v %T %v", j, j, k, k, f, f, g, g)
}
