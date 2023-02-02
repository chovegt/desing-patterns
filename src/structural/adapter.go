package structural

import (
	"fmt"

	"github.com/DioCoding/desing-patterns/src/structural/adapter"
)

func NewAdapter(p adapter.Person) adapter.PersonDTO {
	fmt.Println("start adapter pattern")
	return adapter.ToDto(p)
}
