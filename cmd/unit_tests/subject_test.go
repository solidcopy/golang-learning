package tests

import (
	"fmt"
	"testing"
)

// Testで始まる名前でテスト関数を定義する
func TestGreeting(t *testing.T) {
	greeting := Greeting()
	fmt.Println(greeting)
	// t.Fatalfを実行するとテスト失敗
	// t.Fatalf("test fail")
}
