package main

import "fmt"

type Box struct {
	width, height int
}

func switchLength(b Box) {
	width, height := b.width, b.height

	b.width = height
	b.height = width
}

func switchLengthP(b *Box) {
	width, height := b.width, b.height

	b.width = height
	b.height = width
}

func main () {
	b := Box{width:100, height:200}
	//コピーに対しての操作になる
	switchLength(b)
	fmt.Printf("%+v\n", b) //%+vでフィールド名も表示

	p_b := Box{width:100, height:200}
	switchLengthP(&p_b)
	fmt.Printf("%+v\n", p_b)

	p_bn := &Box{width:100, height:200} //こちらの方がより自然な使い方
	switchLengthP(p_bn)
	fmt.Printf("%+v\n", p_bn)
}
