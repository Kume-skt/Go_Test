package main

import "fmt"

// make([]T, len, cap)　＜＝これでも作れるよ

type Item struct {
	Name   string
	Effect int
}

func main() {
	fmt.Println("=============基本=================")
	basic()
	fmt.Println("=============自動=================")
	autolist()
	fmt.Println("=============for文================")
	var itemNames = []string{"木の棒", "やくそう", "ウマのふん"}

	for i := 0; i < len(itemNames); i++ {
		fmt.Println(itemNames[i])
	}

	for i, v := range itemNames {
		fmt.Println(i, "番目 :", v)
	}
	for v := range itemNames {
		// このｖは番号が入る
		fmt.Println(v)
	}
}

func basic() {
	// 基本記法
	var itemNames [2]string
	itemNames[0] = "木の棒"
	itemNames[1] = "やくそう"

	fmt.Println(itemNames)
	fmt.Println("要素数", len(itemNames))
}

/*
 要素数を自動的に調整してくれる書き方
*/
func autolist() {
	// itemNames := [...]string{"木の棒", "やくそう"}
	itemNames := []string{"木の棒", "やくそう"}

	fmt.Println(itemNames)
	fmt.Println("要素数", len(itemNames))

	// add
	// Goでは配列は宣言時に要素の数が固定されるので要素の追加はできません。
	newItemNames := append(itemNames, "ウマのふん")

	fmt.Println(newItemNames)
	fmt.Println("要素数", len(newItemNames))

	var itemBox []string

	fmt.Println("代入")
	itemBox = newItemNames
	itemBox[2] = "ヤギのふん"

	fmt.Println(itemBox)
	fmt.Println("要素数", len(itemBox))
}
