package main

import "fmt"

type Person struct {
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
