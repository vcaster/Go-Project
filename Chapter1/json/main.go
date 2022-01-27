package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Val   string
	Value string
}

func main() {
	f := Foo{"Values", "Vallys"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)
}
