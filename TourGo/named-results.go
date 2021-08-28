package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main()  {
	fmt.Println(split(17))
}
//関数の戻り値に名前をつけることができる(今回で言うとx,y)
//戻り値に名前をつけた場合returnステイトメントはnakedで良い
//呼び出し元では普通に2つ値が返される
