package main

import (
	"fmt"
	"regexp"

	"github.com/solidcopy/golang-learning/internal/structs"
)

func main() {
	result := function(10)
	fmt.Println(result)

	// 構造体も値渡し
	person := structs.Person{Name: "hattori", Age: 57}
	changeStruct(person)
	fmt.Printf("person.Age: %d\n", person.Age)

	// 複数の戻り値
	dogOrCat, found := multipleReturnValues("monkey dog mouse")
	fmt.Printf("dogOrCat: %v, found: %v\n", dogOrCat, found)

	// 名前付き戻り値
	twice, half := namedReturnValues(10)
	fmt.Printf("twice: %d, half: %d\n", twice, half)

	// 可変個の引数をとる関数
	// 個別の引数として渡したことにする
	args := []int{1, 2, 3, 4}
	sum := valyLengthArguments(args...)
	fmt.Printf("sum: %d\n", sum)
}

func function(argument int) string {
	return fmt.Sprintf("argument: %d", argument)
}

func changeStruct(person structs.Person) {
	person.Age = 18
	fmt.Printf("changeStruct person.Age: %d\n", person.Age)
}

func multipleReturnValues(line string) (string, bool) {
	pattern := regexp.MustCompile(`(dog|cat)`)
	match := pattern.FindStringSubmatch(line)
	if match != nil {
		return match[1], true
	} else {
		return "", false
	}
}

// 短い関数でのみ使用すべき
// returnは必要
func namedReturnValues(n int) (twice, half int) {
	twice = n * 2
	half = n / 2
	return
}

func valyLengthArguments(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
