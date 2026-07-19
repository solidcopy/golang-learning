package main

import "fmt"

func main() {
	numberTypes()
	textTypes()
	booleanType()
}

func numberTypes() {
	// ビット長はシステム依存
	var signedInteger int
	signedInteger = -123
	fmt.Printf("signedInteger: %d\n", signedInteger)

	var unsignedInteger uint
	unsignedInteger = 456
	fmt.Printf("unsignedInteger: %d\n", unsignedInteger)

	// int, uintにビット長を付与した型
	var twoBytes int16
	twoBytes = 0x7fff
	fmt.Printf("twoBytes: %d\n", twoBytes)

	var eightBytes uint64
	eightBytes = 0xffffffffffffffff
	fmt.Printf("eightBytes: %d\n", eightBytes)

	// uint8には別名がある
	var oneByte byte
	oneByte = 0xff
	fmt.Printf("oneByte: %d\n", oneByte)

	// ただのfloatはない
	var singleFloat float32
	singleFloat = 1.23
	fmt.Printf("singleFloat: %f\n", singleFloat)

	var doubleFloat float64
	doubleFloat = 3.21
	fmt.Printf("doubleFloat: %f\n", doubleFloat)

	// 複素数型のcomplex64, complex128もあるが使わないので省略
}

func textTypes() {
	// 文字型のruneはint32の別名
	var aChar rune
	aChar = '山'
	fmt.Printf("aChar: %c\n", aChar)

	var aString string
	aString = "hello"
	fmt.Printf("aString: %s\n", aString)
}

func booleanType() {
	var yes, no bool
	yes = true
	no = false
	fmt.Printf("yes: %v\n", yes)
	fmt.Printf("no: %v\n", no)
}
