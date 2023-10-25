package main

import "fmt"

type Foo struct {
	bar bool
	baz string
	biz int
}

func main() {
	p := &Foo {
		baz: "hello world",
	}
	v := Foo {
		baz: "hello world",
	}

	fmt.Printf("p: %+v\nv: %+v\n", p, v)
}
