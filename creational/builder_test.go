package creational

import (
	"testing"

	"github.com/DioCoding/desing-patterns/creational/builder"
	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {

	type Arg struct {
		materials string
	}

	tests := []struct {
		name string
		arg  Arg
		want string
	}{
		{
			name: "build wood house",
			arg: Arg{
				materials: "wood",
			},
			want: "door: wood, window: wood",
		},
		{
			name: "build plastic house",
			arg: Arg{
				materials: "plastic",
			},
			want: "door: plastic, window: plastic",
		},
	}

	for _, tt := range tests {
		houseBuilder := builder.NewBuilder()
		houseBuilder.SetDoorType(tt.arg.materials)
		houseBuilder.SetWindowType(tt.arg.materials)
		house := houseBuilder.Build()

		assert.Equal(t, tt.want, house.ToString())
	}

}
