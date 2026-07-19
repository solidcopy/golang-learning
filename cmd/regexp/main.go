package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`\d{3}`)

	match := pattern.FindStringSubmatch("123 456 789")
	fmt.Printf("match: %v\n", match)

	match = pattern.FindStringSubmatch("abc")
	if match == nil {
		fmt.Println("no match")
	}

	// 後方参照
	pattern = regexp.MustCompile(`(\d{3})-(\d{4})`)
	match = pattern.FindStringSubmatch("postal code 270-1356")
	fmt.Printf("match: %v\n", match)
}
