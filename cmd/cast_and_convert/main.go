package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 型名を関数のように使ってキャストする
	unsigned := uint(123)
	fmt.Printf("usigned: %d\n", unsigned)

	// 正常にキャストできない
	minus := -123
	fail := uint(minus)
	fmt.Printf("fail: %d\n", fail)

	// これはコンパイルエラー
	// minus := uint(123)

	// 整数 <=> 文字列
	integerToString := strconv.Itoa(12345)
	fmt.Printf("integerToString: %s\n", integerToString)

	stringToInteger, err := strconv.Atoi("12345")
	if err == nil {
		fmt.Printf("stringToInteger: %d\n", stringToInteger)
	}

	// 文字列 <=> バイナリ
	stringToBinary := []byte("string")
	fmt.Printf("stringToBinary: %v\n", stringToBinary)

	binaryToString := string(stringToBinary)
	fmt.Printf("binaryToString: %v\n", binaryToString)

	// 型アサーション
	// インタフェースから実際の型に変換する
	var stringAsAny any = "string as any"
	assertedString := stringAsAny.(string)
	fmt.Printf("assertedString: %v\n", assertedString)

	var integerAsAny any = 12345
	_, castSuccess := integerAsAny.(string)
	fmt.Printf("castSuccess: %v\n", castSuccess)
}
