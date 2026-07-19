package main

import (
	"fmt"

	"github.com/solidcopy/golang-learning/internal/structs"
)

func main() {
	// newはメモリの確保とゼロ詰めだけを行ってポインタを返す
	personPointer := new(structs.Person)
	fmt.Printf("personPointer: %+v\n", *personPointer)

	// makeはスライス、マップ、チャンネルを作成するのに使う
	// これらにnewを使わないのはゼロ詰めではなく初期化するため
}
