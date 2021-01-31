package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Olá Universo")
}

//Greeter exported as symbol
var Greeter greeting
