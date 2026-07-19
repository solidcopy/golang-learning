package main

import "fmt"

func main() {
	defer stopPanic()
	function1()
}

func function1() {
	defer fmt.Println("function1 defer")
	function2()
}

func function2() {
	defer fmt.Println("function2 defer")
	// パニックはプログラムを停止する
	// パニックが発生するとdeferされた関数を呼び出しながらスタックを巻き戻していき、トップに到達すると終了する
	panic("panic occured")
}

func stopPanic() {
	// deferされた関数内でrecoverを呼ぶとパニックを停止してエラー情報を取得できる
	err := recover()
	fmt.Printf("err: %v\n", err)
}

// アプリでpanicを使うケース
// ・バグがなければ到達しないコード
// ・必須の初期化処理の失敗
// ・未実装を表す

// アプリでrecoverを使うケース
// ・goroutineのパニックでプロセスが終了するのを防ぐ
