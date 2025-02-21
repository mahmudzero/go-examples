package main

import "fmt"

func rec[T any](res *T, err *error) {
	fmt.Println("called rec")

	r := recover()
	if r != nil {
		fmt.Println("called rec: panic")
		var t T
		*res = t
		e := fmt.Errorf("panic %+v", r)
		*err = e
	}
}

func canPanic(doPanic bool) (res int, err error) {
	defer rec(&res, &err)

	res = 5
	err = nil
	if doPanic {
		panic("panic at the disco")
	}

	return res, err
}

func canPanicS(doPanic bool) (res string, err error) {
	defer rec(&res, &err)

	res = "hello world"
	err = nil
	if doPanic {
		panic("panic at the disco")
	}

	return res, err
}

func mut(i *int) {
	*i = 55
}

func main() {
	fmt.Println("ints")
	f, e := canPanic(false)
	fmt.Printf("f: %d, e: %+v\n", f, e)

	s, e := canPanic(true)
	fmt.Printf("s: %d, e: %+v\n", s, e)

	fmt.Println("\nstrings")
	fs, e := canPanicS(false)
	fmt.Printf("f: %s, e: %+v\n", fs, e)

	ss, e := canPanicS(true)
	fmt.Printf("s: %s, e: %+v\n", ss, e)

	fmt.Println("\nmut")
	i := 10
	mut(&i)
	fmt.Printf("%d\n", i)
}
