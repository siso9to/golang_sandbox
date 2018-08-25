package main

import (
	"fmt"
)

func main () {
	var i int

	i = 100

	p := &i

	fmt.Printf("Type: %T\n", p)
	fmt.Printf("i = %d\n", *p)

	fmt.Println("---------------")

	pp := &p

	fmt.Printf("Type: %T\n", pp)
	fmt.Printf("i = %d\n", *pp)     // ポインタをデリファレンス
	fmt.Printf("i = %d\n", **pp) // ポインタのポインタをデリファレンス

	fmt.Printf("address = %p\n", pp)      // %pでアドレスを表示できる
}
