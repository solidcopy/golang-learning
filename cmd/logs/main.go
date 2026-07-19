package main

import "fmt"

type Person struct {
	name string
	age  int
}

// fmt.Stringerインタフェースを実装
func (person Person) String() string {
	return fmt.Sprintf("%v (%d)", person.name, person.age)
}

func main() {
	person := Person{"hattori", 57}
	fmt.Printf("person: %v\n", person)
}
