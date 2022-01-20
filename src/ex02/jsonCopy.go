package main

import (
	// "encoding/json"
	"fmt"
)

type Item struct {
	Name   string
	Effect int
}

func main() {
	fmt.Println("＝＝＝＝＝＝＝開始＝＝＝＝＝＝＝")
	// メモリ上のある場所に変数の値が格納
	item1 := Item{
		Name:   "やくそう",
		Effect: 20,
	}

	fmt.Printf("item1の内容 :%+v\n", item1)

	fmt.Println("=============item2===============")
	fmt.Println("50に書き換えた")
	// 同じアドレスを参照している
	item2 := item1
	item2.Effect = 50

	fmt.Printf("item1の内容 :%+v\n", item1)
	fmt.Printf("item2の内容 :%+v\n", item2)

	// アドレスをコピーした？
	item3 := &item1

	fmt.Println("=============item3===============")
	fmt.Printf("item3のAddress :%p\n", item3)
	fmt.Printf("item3の内容 :%+v\n", item3)

	fmt.Println("============================")
	fmt.Println("item3の内容変更をするとitem1も変化する")

	item3.Effect = 10

	fmt.Printf("item1の内容 :%+v\n", item1)
	fmt.Printf("item3の内容 :%+v\n", item3)

}

/*
実行結果

＝＝＝＝＝＝＝開始＝＝＝＝＝＝＝
item1の内容 :{Name:やくそう Effect:20}
=============item2===============
50に書き換えた
item1の内容 :{Name:やくそう Effect:20}
item2の内容 :{Name:やくそう Effect:50}
=============item3===============
item3の内容 :&{Name:やくそう Effect:20}
============================
item3の内容変更をするとitem1も変化する
item1の内容 :{Name:やくそう Effect:10}
item3の内容 :&{Name:やくそう Effect:10}

*/
