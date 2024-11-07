package main

import (
	"errors"
	"fmt"
)

func can_panic(p bool) error {
	if p {
		panic("omg")
	}
	return errors.New("can_panic_error")
}

type PE struct{ Message string }

func (pe *PE) Error() string {
	return fmt.Sprintf("panic error: %s", pe.Message)
}

type Zift struct{}

func caller(p bool) (zift *Zift, err error) {
	zift = nil
	err = errors.New("")

	defer func() {
		var p any = recover()
		if p != nil {
			fmt.Println("!caught panic!")
			err = errors.New("caught panic")
		}
	}()

	err = can_panic(p)
	zift = &Zift{}

	return zift, err
}

func main() {
	var z *Zift
	var e error

	z, e = caller(false)
	fmt.Printf("no panic: r: %+v, error: %+v\n", z, e)
	z, e = caller(true)
	fmt.Printf("   panic: r: %+v, error: %+v\n", z, e)
}
