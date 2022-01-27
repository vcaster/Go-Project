package main

import (
	"fmt"
)

func main() {
	// // Pointers
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)
}
