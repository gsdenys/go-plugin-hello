package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}

//Greeter exported as symbol
var Greeter greeting
