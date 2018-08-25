package main

import "fmt"

type ColorBox struct {
	width, height int
}

func NewColorBox(width int, height int) *ColorBox {
	b := new(ColorBox)
	b.width = width
	b.height = height

	return b
}

func main() {
	nb := new(ColorBox)
	fmt.Printf("%+v\n", nb)

	nb.width = 100
	nb.height = 150
	fmt.Printf("%+v\n", nb)

	fmt.Println("--address operator--")

	nb_address := &ColorBox{}
	fmt.Printf("%+v\n", nb_address)

	nb_address2 := &ColorBox{width: 100, height: 150}
	fmt.Printf("%+v\n", nb_address2)

	fmt.Println("--constructor--")

	nc := NewColorBox(400, 500)
	fmt.Printf("%+v\n", nc)
}
