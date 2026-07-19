package main

import "fmt"

func main() {
	n := 123
	var p *int
	p = &n
	fmt.Printf("*p: %d\n", *p)

	*p = 456
	fmt.Printf("*p: %d\n", *p)

	// ポインタ演算はない
}
