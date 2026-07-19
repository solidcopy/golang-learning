package main

import "fmt"

// 初期化処理を行う関数
// 引数、戻り値はなし
func init() {
	fmt.Println("init")
}

// initは複数定義できる
func init() {
	fmt.Println("more init")
}

func main() {
	fmt.Println("main")
}
