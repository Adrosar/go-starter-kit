package foo

// Foo ...
type Foo struct {
	name string
	age  int
}

// NewFoo ...
func NewFoo(name string) *Foo {
	return &Foo{name: name}
}

// SetAge ...
func (foo *Foo) SetAge(age int) {
	foo.age = age
}

// GetAge ...
func (foo *Foo) GetAge() int {
	return foo.age
}

// GetName ...
func (foo *Foo) GetName() string {
	return foo.name
}
