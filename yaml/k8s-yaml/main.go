package main

import (
	"fmt"

	yv3 "sigs.k8s.io/yaml"
)

var data = `
k1: 5
k3: 50
`

type Foo struct {
	K1 int  `yaml:"k1"`
	K2 *int `yaml:"k2"`
	K3 int  `default:"100" yaml:"k3"`
}

func main() {
	var f Foo
	var err error

	err = yv3.Unmarshal([]byte(data), &f)
	if err != nil { panic(err) }

	fmt.Printf("%+v\n", f)
}
