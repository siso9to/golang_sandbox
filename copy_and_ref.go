package main

import "fmt"

func main() {
	var value int

	value = 1

	//値をそのまま渡す
	fmt.Printf("before value: %d\n", value)
	copy_value(value)
	fmt.Printf("after value: %d\n", value)

	p_value := &value

	//値のポインタを渡す
	fmt.Printf("before p_value: %d\n", *p_value)
	ref_value(p_value)
	fmt.Printf("after p_value: %d\n", *p_value)

}


func copy_value (value int) {
	value = 10
}

func ref_value (p_value *int) {
	*p_value = 10
}