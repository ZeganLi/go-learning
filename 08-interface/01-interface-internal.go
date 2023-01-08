package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBed = MyError{errors.New("bed error")}

func bed() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bed() {
		p = &ErrBed
	}
	return p
}

func main() {

	e := returnsError()

	if e != nil {
		fmt.Printf("error: %+v\n", e)
		return
	}

	fmt.Println("ok")
}
