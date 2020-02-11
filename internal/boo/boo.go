package boo

// Boo ...
type Boo struct {
	name string
	age  int
}

// NewBoo ...
func NewBoo(name string) *Boo {
	return &Boo{name: name}
}

// SetAge ...
func (boo *Boo) SetAge(age int) {
	boo.age = age
}

// GetAge ...
func (boo *Boo) GetAge() int {
	return boo.age
}

// GetName ...
func (boo *Boo) GetName() string {
	return boo.name
}
