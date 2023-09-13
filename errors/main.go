package main

import (
	"fmt"
	"main.com/cerrors"
)

func MakeErr() error {
	return &cerrors.CErr { Msg: "I am an error" }
}

func main() {
	err := MakeErr()
	if _, ok := err.(*cerrors.CErr); ok {
		fmt.Println("The type is also good")
	}
	fmt.Printf("err: %+v\n", err)
}
