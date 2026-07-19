package main

import (
	"fmt"
)

func main() {
	score := 50

	if score >= 80 {
		fmt.Println("very good")
	} else if score >= 50 {
		fmt.Println("good")
	} else {
		fmt.Println("poor")
	}

	// ifだけで有効な変数
	price := 1000
	quantity := 200
	if total := price * quantity; total >= 100000 {
		fmt.Printf("Too expensive. total: %d\n", total)
	}

	// switchはフォールスルーしないのでbreakは不要
	highScore := 80
	switch {
	case highScore >= 80:
		fmt.Println("very good")
	case highScore >= 50:
		fmt.Println("good")
	default:
		fmt.Println("poor")
	}

	// 型による分岐
	var scoreAsAny any = score
	switch scoreAsAny.(type) {
	case int:
		fmt.Println("score is int")
	default:
		fmt.Println("never")
	}
}
