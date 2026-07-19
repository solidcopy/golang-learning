package main

import "fmt"

func main() {
	numberMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("numberMap[\"two\"]: %d\n", numberMap["two"])

	numberMap["four"] = 4
	fmt.Printf("numberMap[\"four\"]: %d\n", numberMap["four"])

	delete(numberMap, "three")
	// 存在しないキーを指定すると、値型のゼロ値を返す
	fmt.Printf("numberMap[\"three\"]: %d\n", numberMap["three"])

	// キーの存在確認
	_, hasOne := numberMap["one"]
	_, hasThree := numberMap["three"]
	fmt.Printf("hasOne: %v, hasThree: %v\n", hasOne, hasThree)

	// TODO: makeを使うとどういう違いがある？
	dynamicMap := make(map[int]string)
	dynamicMap[4] = "four"
	fmt.Printf("dynamicMap[4]: %v\n", dynamicMap[4])
}
