package behavioral

import (
	"fmt"

	"github.com/DioCoding/desing-patterns/behavioral/strategies"
)

func ExecuteStrategyPattern() {
	fmt.Println("start strategy pattern")
	text := "Hello strategy :D"

	// with structs
	fmt.Println("Cloud strategy")
	cloudStruct := strategies.New(text, strategies.SaveCloud{})
	fmt.Println(cloudStruct.Print())

	fmt.Println("Local strategy")
	localStruct := strategies.New(text, strategies.SaveLocal{})
	fmt.Println(localStruct.Print())

	// with functions
	fmt.Println("Cloud strategy")
	fCloud, err := strategies.Strategy("cloud")
	if err == nil {
		fmt.Println(fCloud(text))
	}

	fmt.Println("Local strategy")
	fLocal, err := strategies.Strategy("local")
	if err == nil {
		fmt.Println(fLocal(text))
	}

	fmt.Println("invalid strategy")
	_, err = strategies.Strategy("invalid")
	fmt.Println(err)
}
