package main

import (
	"fmt"

	"github.com/solidcopy/golang-learning/internal/structs"
)

type MyPerson structs.Person

// fmt.Stringerを実装して%vの出力形式を定義する
func (p MyPerson) String() string {
	return fmt.Sprintf("%v (%d)", p.Name, p.Age)
}

func main() {
	n := 123456
	fmt.Println("n := 10000")
	fmt.Printf("%%d: %d\n", n)
	fmt.Printf("%%b: %b\n", n)
	fmt.Printf("%%o: %o\n", n)
	fmt.Printf("%%O: %O\n", n)
	fmt.Printf("%%x: %x\n", n)
	fmt.Printf("%%X: %X\n", n)
	fmt.Println()

	f := 12.3456
	fmt.Println("f := 12.3456")
	fmt.Printf("%%f: %f\n", f)
	// 他にもあるが使いそうにないので省略
	fmt.Println()

	b := true
	fmt.Println("b := true")
	fmt.Printf("%%t: %t\n", b)
	fmt.Println()

	c := '水'
	fmt.Println("c := '水'")
	fmt.Printf("%%c: %c\n", c)
	fmt.Printf("%%q: %q\n", c)
	fmt.Printf("%%U: %U\n", c)
	fmt.Printf("%%d: %d\n", c)
	fmt.Println()

	s := "hello"
	fmt.Println("s := \"hello\"")
	fmt.Printf("%%s: %s\n", s)
	fmt.Printf("%%q: %q\n", s)
	fmt.Printf("%%x: %x\n", s)
	fmt.Printf("%%X: %X\n", s)
	fmt.Println()

	person := structs.Person{Name: "hattori", Age: 57}
	fmt.Println("person := structs.Person{name: \"hattori\", age: 57}")
	fmt.Printf("%%v: %v\n", person)
	fmt.Printf("%%+v: %+v\n", person)
	fmt.Printf("%%#v: %#v\n", person)
	fmt.Printf("%%T: %T\n", person)
	fmt.Println()

	pointer := &person
	fmt.Println("pointer := &person")
	fmt.Printf("%%p: %p\n", pointer) // ポインタのアドレス
	fmt.Println()

	myPerson := MyPerson{Name: "hattori", Age: 57}
	fmt.Println("myPerson := MyPerson{name: \"hattori\", age: 57}")
	fmt.Printf("%%v: %v\n", myPerson)
}
