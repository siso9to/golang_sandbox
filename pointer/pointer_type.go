package main

import "fmt"

func main() {
	// int型のポインタを定義
	var hoge *int

	*hoge = 100
	fmt.Printf("%d", *hoge) //エラー
}
