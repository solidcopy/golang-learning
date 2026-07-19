package main

import "fmt"

func main() {
	// 初期化していない変数のゼロ値

	// 0
	var integer int
	fmt.Printf("int: %d\n", integer)

	var float float64
	fmt.Printf("float64: %f\n", float)

	// false
	var boolean bool
	fmt.Printf("bool: %v\n", boolean)

	// ''
	var char rune
	fmt.Printf("rune: '%c'\n", char)

	// ""
	var stringVar string
	fmt.Printf("string: \"%s\"\n", stringVar)

	// nil
	var slice []string
	fmt.Printf("slice: %v, slice is nil: %v\n", slice, slice == nil)
}
