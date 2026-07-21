package main

import "fmt"

func main() {
	// 無限ループ
	// break, continue
	n := 0
	fmt.Print("n: ")
	for {
		n++

		if n >= 10 {
			break
		} else if n%2 == 1 {
			continue
		}

		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// while
	n = 10
	fmt.Print("n: ")
	for n >= 7 {
		fmt.Printf("%d ", n)
		n -= 1
	}
	fmt.Println()

	// スライスのイテレート
	slice := []string{"apple", "banana", "cinamon"}
	for i, element := range slice {
		fmt.Printf("slice[%d]: %v\n", i, element)
	}

	// ラベル付きもあるが使わないので省略

	// マップのイテレート
	// 順序は保証されていない
	numberMap := make(map[string]int)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	for k, v := range numberMap {
		fmt.Printf("key: %v, value: %d\n", k, v)
	}
}
