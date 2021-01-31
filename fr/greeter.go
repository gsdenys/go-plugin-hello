package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Bonjour Univers")
}

//Greeter exported as symbol
var Greeter greeting
