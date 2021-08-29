package main

import "fmt"

func main()  {
	defer fmt.Println("World")
	defer fmt.Println("unti")
	fmt.Println("Hello")
}
//deferステートメントはdeferへ渡した関数の実行を呼び出し元の関数の終わり（return）するまで延期させるもの
//deferへ渡した関数の引数はすぐに評価される
//deferが複数ある場合はLIFOの順番でstackされる
