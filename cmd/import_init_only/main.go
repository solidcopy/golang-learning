package main

// init関数だけ実行する
import (
	"fmt"

	_ "github.com/solidcopy/golang-learning/cmd/import_init_only/init"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
