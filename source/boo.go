package main

func main() {
	foo1 := createFoo("Foo1")
	println("ID:", foo1.ID)
	println("Name:", foo1.Name)

	foo2 := createFoo("Foo2")
	println("ID:", foo2.ID)
	println("Name:", foo2.Name)

	foo3 := createFoo("Foo3")
	println("ID:", foo3.ID)
	println("Name:", foo3.Name)
}
