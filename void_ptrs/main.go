package main

import (
	"fmt"
	"errors"
)

type Foo struct {
	Name string
	Age  int
}

func ErrAndIfc(ipt interface{}) (interface{}, error) {
	return ipt, errors.New("tmp");
}

func main() {
	var str string = "hello world";
	var num int    = 50;
	var foo Foo    = Foo {
		Name: "Johnny",
		Age:  25,
	};

	var retStr string;
	var retNum int;
	var retFoo Foo;
	var interm interface{}

	var err error;

	interm, err = ErrAndIfc(str);
	retStr = interm.(string);
	fmt.Printf("str: %s, err: %+v\n", retStr, err);

	interm, err = ErrAndIfc(num);
	retNum = interm.(int);
	fmt.Printf("num: %d, err: %+v\n", retNum, err);

	interm, err = ErrAndIfc(foo);
	retFoo = interm.(Foo);
	fmt.Printf("foo: %+v, err: %+v\n", retFoo, err);
}
