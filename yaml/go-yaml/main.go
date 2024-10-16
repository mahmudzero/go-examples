package main

import (
	"fmt"

	yv3 "gopkg.in/yaml.v3"
)

var data = `
k1: 5
`

type Foo struct {
	K1 int  `yaml:"k1"`
	K2 *int `yaml:"k2"`
	K3 int  `yaml:"k3" default:"100"`
}

func main() {
	var f Foo
	var err error

	err = yv3.Unmarshal([]byte(data), &f)
	if err != nil { panic(err) }

	fmt.Printf("%+v\n", f)
}
