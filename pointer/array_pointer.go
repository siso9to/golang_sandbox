package main

import (
	"fmt"
)

func main() {
	v := [3]int{1, 2, 3}
	fmt.Printf("%d\n", v[0]) // これは普通

	p := &[3]int{1, 2, 3}
	//fmt.Printf("%d\n", *p[0]) // コンパイルエラー
	fmt.Printf("%d\n", (*p)[0]) // デリファレンス
	fmt.Printf("%d\n", p[0]) // *つけなくてもデリファレンスできてる


	fmt.Println("built in functions and slice")
	fmt.Println(len(p))
	fmt.Println(cap(p))
	fmt.Println(p[0:2])

	//builtInFunc(p)
}

func builtInFunc(array *[3]int) {
	fmt.Println("built in functions and slice")
	fmt.Println(len(array))
	fmt.Println(cap(array))
	fmt.Println(array[0:2])
}
