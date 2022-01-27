package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("Hello : ", p.Name)
}

func main() {
	var indiv = new(Person)
	indiv.Name = "Zeke"
	indiv.SayHello()
}
