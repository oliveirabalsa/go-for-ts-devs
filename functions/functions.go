package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println("Hello World")
	withParams("Leonardo")
	fmt.Println(oneReturnedValue("Leonardo"))
	fmt.Println(twoReturnedValues("Leonardo"))
	name, nameLength := twoReturnedValues("Leonardo")
	fmt.Println(name, nameLength)
	fmt.Println(withTypeReturn("Leonardo"))
	fmt.Println(namedReturnValues("Leonardo"))
}

func withParams(name string) {
	fmt.Println("Hello", name)
}

func oneReturnedValue(name string) string {
	return "Hello " + name
}

func twoReturnedValues(name string) (string, int) {
	return name, len(name)
}

func namedReturnValues(name string) (returnedName string, nameLength int) {
	returnedName = name
	nameLength = len(name)
	return
}

func withTypeReturn(name string) User {
	return User{Name: name, Age: 50, Email: "oliveirabalsa2@gmail.com"}
}
