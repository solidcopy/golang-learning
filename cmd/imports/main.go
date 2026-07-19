package main

import (
	"fmt"

	"github.com/solidcopy/golang-learning/cmd/import/another_module"
)

func main() {
	returnValue := another_module.PublicFunction()
	fmt.Println(returnValue)
}
