package main

type Foo struct {
	ID   int
	Name string
}

var currentID int = 0

func createFoo(name string) *Foo {
	currentID = currentID + 1

	return &Foo{
		ID:   currentID,
		Name: name,
	}
}
