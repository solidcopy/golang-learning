package another_module

import "fmt"

// 大文字から始まる名前は他のパッケージから参照できる
func PublicFunction() string {
	number := privateFunction()
	return fmt.Sprintf("number: %d", number)
}

func privateFunction() int {
	return 123
}
