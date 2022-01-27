package main

import (
	"fmt"
)

func main() {
	// Slice
	var s = make([]string, 0)
	//Map
	var m = make(map[string]string)
	s = append(s, "some string")
	m["some key"] = "some value"
	fmt.Println(s, m)
}
