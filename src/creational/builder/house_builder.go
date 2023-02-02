package builder

import "fmt"

type House struct {
	doorType   string
	windowType string
}

func (h House) ToString() string {
	return fmt.Sprintf("door: %v, window: %v", h.doorType, h.windowType)
}

type HouseBuilder struct {
	windowType string
	doorType   string
}

func NewBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) SetWindowType(window string) {
	b.windowType = window
}

func (b *HouseBuilder) SetDoorType(door string) {
	b.doorType = door
}

func (b *HouseBuilder) Build() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
	}
}
