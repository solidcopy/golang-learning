package main

import "fmt"

// iotaは0からの連番
// 1からにしたければ先頭の定数名を'_'にする。
// 同じconstにまとめた値のない定数は前の式と同じになる
const (
	success = iota
	warning
	error
)

func main() {
	fmt.Printf("success: %d\n", success)
	fmt.Printf("warning: %d\n", warning)
	fmt.Printf("error: %d\n", error)
}
