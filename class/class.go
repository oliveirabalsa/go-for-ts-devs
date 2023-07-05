package main

import "fmt"

type Person struct {
	// !INFO: In Go, capitalizing the field names in a struct indicates that the fields are public, meaning they can be accessed from outside the package.
	//ex: 
	//Name string  ---- means Name is public
	//age int      ---- and age is private
	name string
	age  int
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) setAge(age int) {
	p.age = age
}

func (p *Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func main() {
	// Class instance
	person := Person{name: "Bianca", age: 25}
	person.greet()

	// Accessing class methods
	name := person.getName()
	age := person.getAge()
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	person.setName("Leo")
	person.setAge(30)
	person.greet()
}
