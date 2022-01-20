package main

import (
	// "encoding/json"
	"fmt"
)

type Test struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	fmt.Println("＝＝＝＝＝＝＝開始＝＝＝＝＝＝＝")
	item1 := &Test{
		Name: "kumemaru",
		Age:  20,
	}

	fmt.Println(item1)
	// Value
	fmt.Printf("item1 :%+v\n", item1)
	// Addressポイント
	fmt.Printf("item1 :%p", item1)

}
