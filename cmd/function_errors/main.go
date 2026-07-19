package main

// 複数インポートはこの形
import (
	"errors"
	"fmt"
)

func main() {
	var result string
	var err error

	result, err = function("my value")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Printf("error: %v\n", err)
	}

	result, err = function("")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

func function(argument string) (string, error) {
	if argument == "" {
		return "", errors.New("argument is nil")
	}

	return fmt.Sprintf("argument: %v", argument), nil
}
