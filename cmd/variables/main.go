package main

import "fmt"

func main() {
	// 宣言してから初期化
	var number int
	number = 10
	fmt.Printf("number: %d\n", number)

	// 宣言と同時に初期化
	// 型は初期値で判断される
	greeting := "hello"
	fmt.Printf("greeting: %v\n", greeting)

	// 関数型の変数
	var function func(int) int
	// クロージャにする
	a := 10
	function = func(b int) int {
		return a + b
	}
	fmt.Printf("function(10): %d\n", function(10))
}
