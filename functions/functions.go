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

func withTypeReturn(name string) User {
	return User{Name: name, Age: 50, Email: "oliveirabalsa2@gmail.com"}
}

func operators(x interface{}, y int) bool {
	if x.(int) > y {
		return true
	}

	if x.(int) > 1 || y > 1 {
		return true
	}

	if x.(int) > 1 && y > 1 {
		return true
	}

	if x.(int) != 1 {
		return false
	}

	return false
}
