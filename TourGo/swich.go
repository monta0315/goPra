package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linex .")
	default:
		fmt.Printf("%s.\n",os)
	}
}

//switchはif-elseをきれいに描くためにある〜
//breakステートメントは自動的に提供される
//defaultはcaseを抜けたら
