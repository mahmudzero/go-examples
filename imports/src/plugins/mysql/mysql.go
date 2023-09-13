package mysql

import (
	"fmt"
	"imports.com/main/src/plugins"
)

func DoSomething() {
	fmt.Printf("struct: %+v\n", plugins.Plugin{ Name: "Johnny" })
}
