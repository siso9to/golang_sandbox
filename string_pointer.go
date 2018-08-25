package main

import "fmt"

func main() {
	s := "HelloCamp!"

	fmt.Printf("%T\n", &s)
	fmt.Printf("%T\n", s[0])
	fmt.Printf("%s\n", &s[0]) // cannot take the address of s[0]
}

