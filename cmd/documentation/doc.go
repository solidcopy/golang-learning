// $ go doc
// でドキュメントを出力する。
//
// パッケージ前のコメントはパッケージコメントになる。
//
// 改行は
// 空白に置き換えられる。
//
// 空白行を挟むと段落が分かれる。
//
// 複数の.goがある時はどれに書いてもいい。
package doc

import "fmt"

func SomeFunction(s string) int {
	fmt.Println(s)
	return 123
}
