package main

import (
    // "encoding/json"
    "fmt"
)

type Test struct {
	name string `json:"Name"`
	age int `json:"Age"`
}

func main() {
    fmt.Println("＝＝＝＝＝＝＝開始＝＝＝＝＝＝＝")
    item1 := &Test{
        name: "粂田",
        age:  20,
    }

    fmt.Println(item1)
	// Value
	fmt.Printf("item1 :%+v\n", item1)
	// Addressポイント
	fmt.Printf("item1 :%p", item1)

}