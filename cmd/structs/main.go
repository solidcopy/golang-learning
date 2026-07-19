package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}

	// 構造体リテラルでは全フィールドを書く
	point := Point{56, 67}
	fmt.Printf("point x: %d, y: %d\n", point.X, point.Y)
	point.X = 65
	fmt.Printf("point x: %d, y: %d\n", point.X, point.Y)

	// p.x で (*p).x と同じ
	pointPointer := &point
	fmt.Printf("pointPointer.x: %d\n", pointPointer.X)

	// 1つもフィールドを書かなければすべてゼロ値
	zeroPoint := Point{}
	fmt.Printf("zeroPoint: %+v\n", zeroPoint)

	// フィールド名付きなら一部フィールドを省略してゼロ値にできる
	onlyY := Point{Y: 123}
	fmt.Printf("onlyY: %+v\n", onlyY)

	// 型名なし
	nonameStruct := struct {
		name string
		age  int
	}{"hattori", 57}
	fmt.Printf("nonameStruct: %+v\n", nonameStruct)

	// 型名なしの配列
	nonameStructArray := []struct {
		item  string
		price int
	}{{"Soap", 6}, {"Drink", 2}}
	fmt.Printf("nonameStructArray: %+v\n", nonameStructArray)

	// 構造体のフィールドとして別の構造体を埋め込むと、
	// 内部の構造体のフィールドを自身のそれのように扱える
	type Person struct {
		name string
		age  int
	}
	type Address struct {
		prefecture string
		city       string
	}
	type Profile struct {
		Person
		Address
	}
	profile := Profile{Person{name: "hattori", age: 57}, Address{prefecture: "Saitama", city: "Kasukabe"}}
	fmt.Printf("profile.name: %v\n", profile.name)
	fmt.Printf("profile.Address: %+v\n", profile.Address)
	fmt.Printf("profile: %+v\n", profile)

	// 埋め込むのがポインタなら使う前に実体のnewが必要
}
