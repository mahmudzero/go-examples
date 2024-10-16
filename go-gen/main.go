package main

import "fmt"

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml metrics-oapi.json

func main() {
	fmt.Println("hello world");
}
