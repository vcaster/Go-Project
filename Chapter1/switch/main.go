package main

import (
	"fmt"
)

func main() {
	x := "bar"
	switch x {
	case "foo":
		fmt.Println("I am foo")
	case "bar":
		fmt.Println("I am bar")
	default:
		fmt.Println("Default case")
	}

}
