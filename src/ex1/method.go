package main

import "fmt"

func main(){
	fmt.Println(sum(5))
	printer(5)
}

func sum(q int) int {
	return q+10
}

func printer(q int) {
	fmt.Println(q * 10)
}