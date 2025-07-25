package main

// Objects -------------------------------------------------------------------------------------------------------------
// IBaseObject
type IBaseObject interface {
	ID() string
}

// BaseObject
type BaseObject struct{}

func (*BaseObject) ID() string { return "BaseObject" }

// IComplexObject
type IComplexObject interface {
	IBaseObject
	X() string
}

// ComplexObject
type ComplexObject struct{}

func (*ComplexObject) ID() string { return "ComplexObject" }
func (*ComplexObject) X() string  { return "X" }

// Objects -------------------------------------------------------------------------------------------------------------

// Getters -------------------------------------------------------------------------------------------------------------
type QueryProvider interface {
	GetAll(query string) IBaseObject
}

type DB struct{}

func (*DB) GetAll[T any](query string) IBaseObject {
	switch query {
	case "base":
		return &BaseObject{}
	case "complex":
		return &ComplexObject{}
	default:
		return nil
	}
}

// Getters -------------------------------------------------------------------------------------------------------------

func main() {}
