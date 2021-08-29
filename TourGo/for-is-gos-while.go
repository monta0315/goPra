package main

import "fmt"

func main()  {
	sum := 1
	for sum < 1000{
		sum+=sum
	}
	fmt.Println(sum)
}
//goにwhileはなくfor文に初期化ステートメントと条件をつけないで実装する
