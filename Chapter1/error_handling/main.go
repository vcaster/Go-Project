package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func foo() error {
	return errors.New("some error occurred")
}
func main() {
	if err := foo(); err != nil {
		fmt.Println("Error please")
	}
}
