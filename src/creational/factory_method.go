package creational

import (
	"errors"
	"fmt"

	"github.com/DioCoding/desing-patterns/src/creational/factorymethod"
)

var factories = map[string]factorymethod.Guns{
	"AK47":     factorymethod.NewAk47(),
	"revolver": factorymethod.NewRevolver(),
}

/*
Factory Method is a creational design pattern that provides an interface for creating objects in a superclass,
but allows subclasses to alter the type of objects that will be created.
*/
func NewFactoryMethod(typeGun string) (factorymethod.Guns, error) {
	fmt.Println("start factory method")

	if gun, ok := factories[typeGun]; ok {
		return gun, nil
	}

	return nil, errors.New("invalid type gun")
}
