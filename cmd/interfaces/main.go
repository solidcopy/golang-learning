package main

import "fmt"

type Barker interface {
	Bark() string
}

type Walker interface {
	Walk()
}

// インタフェースの合成
type Animal interface {
	Barker
	Walker
}

// implementsとは書かない
type Dog struct{}

func (dog Dog) Bark() string {
	return "Bow wow"
}

func (dog Dog) Walk() {
	fmt.Println("dog walk")
}

type Cat struct{}

func (cat Cat) Bark() string {
	return "Meow"
}

func (cat Cat) Walk() {
	fmt.Println("cat walk")
}

func main() {
	var animal Animal

	animal = Dog{}
	fmt.Printf("animal.Bark(): %v\n", animal.Bark())
	animal.Walk()

	animal = Cat{}
	fmt.Printf("animal.Bark(): %v\n", animal.Bark())
	animal.Walk()
}
