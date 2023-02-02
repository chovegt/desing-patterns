package creational

import (
	"fmt"

	"github.com/DioCoding/desing-patterns/src/creational/builder"
)

func ExecuteBuilderPattern() {
	fmt.Println("start builder pattern")

	houseBuilder := builder.NewBuilder()
	houseBuilder.SetDoorType("wood")
	houseBuilder.SetWindowType("wood")
	house := houseBuilder.Build()

	fmt.Println(house.ToString())
}
