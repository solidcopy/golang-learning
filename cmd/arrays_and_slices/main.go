package main

import "fmt"

func main() {
	// スライス自体はデータを保持して折らず配列を参照している

	stringSlice := []string{"apple", "banana", "cinamon"}
	fmt.Printf("stringSlice[0]: %v\n", stringSlice[0])
	fmt.Printf("stringSlice[2]: %v\n", stringSlice[2])
	// これはできない
	// stringSlice[-1]

	fmt.Printf("stringSlice[0:2]: %v\n", stringSlice[0:2])
	fmt.Printf("stringSlice[1:]: %v\n", stringSlice[1:])
	fmt.Printf("stringSlice[:2]: %v\n", stringSlice[:2])

	stringTable := [][]string{{"name", "age"}, {"hattori", "57"}}
	fmt.Printf("stringTable[1][0]: %v\n", stringTable[1][0])

	length := len(stringSlice)
	fmt.Printf("length: %d\n", length)

	// 長いスライスから短いスライスを切り出す
	longSlice := []int{1, 2, 3, 4, 5, 6}
	shortSlice := longSlice[:3]
	fmt.Printf("len(shortSlice): %d, cap(shortSlice): %d\n", len(shortSlice), cap(shortSlice))
	// capを上限としてより長いスライスにすることもできる
	extendedSlice := shortSlice[:5]
	fmt.Printf("extendedSlice: %v\n", extendedSlice)

	// インデックスを指定して初期化
	onlyOdds := []int{1: 2, 3: 6, 5: 10}
	fmt.Printf("onlyOdds[3]: %d\n", onlyOdds[3])
	fmt.Printf("onlyOdds[4]: %d\n", onlyOdds[4])
	fmt.Printf("len(onlyOdds): %d\n", len(onlyOdds))

	// 要素数が事前に分かっているならmakeを使う
	// 指定した容量までならappendで再割り当てが発生しない
	sliceSize := 5
	dynamicAllocated := make([]int, sliceSize)
	fmt.Printf("len(dynamicAllocated): %d\n", len(dynamicAllocated))

	// スライスへの追加
	// 容量が足りていれば同じスライスに要素を追加するが、足りなければ新しいスライスになる
	stringSlice = append(stringSlice, "durian", "eggplant")
	fmt.Printf("stringSlice: %v\n", stringSlice)

	// 要素数を明示すると配列となり、動的なサイズ変更(append)はできない
	// 値渡しでは参照ではなく配列がコピーされる
	stringArray := [2]string{"blue", "red"}
	fmt.Printf("stringArray: %v\n", stringArray)
	// これでも同じ
	// stringArray := [...]string{"blue", "red"}
}
