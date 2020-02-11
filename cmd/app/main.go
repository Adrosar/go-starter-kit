package main

import (
	"fmt"
	internalBoo "go-starter-kit/internal/boo"
	pkgFoo "go-starter-kit/pkg/foo"
)

func main() {
	boo := internalBoo.NewBoo("Boo")
	fmt.Println("[1] Boo:", boo)
	boo.SetAge(12)
	fmt.Println("[2] Boo:", boo)
	fmt.Println("[3] Name:", boo.GetName())
	fmt.Println("[4] Age:", boo.GetAge())

	foo := pkgFoo.NewFoo("Foo")
	fmt.Println("[5] Foo name:", foo.GetName())
	fmt.Println("[5] Foo age:", foo.GetAge())
}
