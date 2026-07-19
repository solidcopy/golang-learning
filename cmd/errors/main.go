package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	errorCode    string
	errorMessage string
}

func (err MyError) Error() string {
	return fmt.Sprintf("[%v]: %v", err.errorCode, err.errorMessage)
}

func main() {
	answer, err := sayYes("yes")
	showResult(answer, err)

	answer, err = sayYes("no")
	showResult(answer, err)

	showResult("", MyError{errorCode: "E0001", errorMessage: "Internal Error"})
}

// errorは組み込みのインタフェース
// Error() string
func sayYes(message string) (string, error) {
	if message == "yes" {
		return "good", nil
	} else {
		return "", errors.New("bad")
	}
}

func showResult(answer string, err error) {
	if err == nil {
		fmt.Printf("success! answer: %v\n", answer)
	} else {
		fmt.Printf("error! err.Error(): %v\n", err.Error())
	}
}
