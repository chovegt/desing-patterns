package main

import (
	"fmt"

	"github.com/DioCoding/desing-patterns/src/behavioral"
)

func main() {
	saludo := "** Learn with Dio **\nHappy coding"
	fmt.Println(saludo)

	// fmt.Println("Creationals")
	// creational.ExecuteBuilderPattern()
	// creational.NewFactoryMethod()

	fmt.Println("Behaviorals")
	behavioral.ExecuteStrategyPattern()

	// fmt.Println("Structurals")
	// structural.NewAdapter()

	fmt.Println("Bye, vuelvan prontus!!")
}
