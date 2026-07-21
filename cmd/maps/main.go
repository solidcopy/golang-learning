package main

import (
	"fmt"
	"sync"
)

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

	// 事前にエントリ数が分かっているならmakeに渡して最適化のヒントにできる
	dynamicMap := make(map[int]string, 100)
	dynamicMap[4] = "four"
	fmt.Printf("dynamicMap[4]: %v\n", dynamicMap[4])

	// マップのキーに使えるのは、
	// ・数値
	// 文字列
	// ポインタ
	// チャネル
	// インタフェース
	// それらのみを含む配列と構造体
	// スライス、マップ、関数は不可

	// マップはマルチスレッドセーフではないので
	// 並列で読み書きするなら排他制御する
	synchronizedMap := struct {
		sync.RWMutex
		m map[int]string
	}{m: make(map[int]string)}

	// 書き込みロック
	synchronizedMap.Lock()
	synchronizedMap.m[11] = "eleven"
	synchronizedMap.Unlock()

	// 読み取りロック
	synchronizedMap.RLock()
	value := synchronizedMap.m[11]
	synchronizedMap.RUnlock()
	fmt.Printf("value: %v\n", value)
}
