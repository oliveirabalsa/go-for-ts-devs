package main

import (
	"fmt"
	"reflect"
)

func main() {
	// !INFO: declarar uma variavel
	var x string
	fmt.Println("String: ", x)
	fmt.Println(reflect.TypeOf(x))

	//  var x string = "Hello World" // Remove o var e remove o tipo para inicializar
	// fmt.Println("String: ", x)

	// !INFO: declarar e inicializar
	var y string = "Hello World"
	fmt.Println(y)

	// !INFO: declarar e inicializar sem o tipo (inferencia de tipo)
	var z = "Hello World"
	fmt.Println(z)

	// !INFO: declarar e inicializar sem o tipo e sem o var
	w := "Hello World"
	fmt.Println(w)

	// !INFO: inferiu o tipo float64
	pi := 3.14
	fmt.Println(pi)
	fmt.Println(reflect.TypeOf(pi))

	const PI = 3.1415
	// !INFO: tentativa de trocar a variavel const
	// PI = 1

	calc := pi * 2
	fmt.Println(calc)
}
