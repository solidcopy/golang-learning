package main

import "fmt"

func main() {
	hello()
	multipleDefer()
}

func hello() {
	// 関数実行をdeferすると引数は評価されるが実行は関数の終了時まで遅延する
	// finallyのような使い方をする
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func multipleDefer() {
	// 複数回deferするとスタックして逆順に実行する
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	fmt.Println("One")
}
