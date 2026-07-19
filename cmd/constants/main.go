package main

import "fmt"

// 定数名の付け方は変数と同じ
// キャメルケースで、公開するなら頭を大文字にする

// 以下の型のみ可
const pi = 3.14

const score = 100

const rating = 'S'

const greeting = "hello"

const available = true

func main() {
	fmt.Printf("pi: %f\n", pi)
	fmt.Printf("score: %d\n", score)
	fmt.Printf("rating: %c\n", rating)
	fmt.Printf("greeting: %v\n", greeting)
	fmt.Printf("available: %v\n", available)
}
