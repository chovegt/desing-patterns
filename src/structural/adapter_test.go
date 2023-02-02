package structural_test

import (
	"testing"

	"github.com/DioCoding/desing-patterns/src/structural"
	"github.com/DioCoding/desing-patterns/src/structural/adapter"
	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	tests := []struct {
		name string
		args adapter.Person
		want string
	}{
		{
			name: "Pedro Lopez",
			args: adapter.Person{
				Name:     "Pedro",
				LastName: "Lopez",
			},
			want: "Pedro Lopez",
		},
		{
			name: "Firulay Rodriguez",
			args: adapter.Person{
				Name:     "Firulay",
				LastName: "Rodriguez",
			},
			want: "Firulay Rodriguez",
		},
	}

	for _, tt := range tests {
		personDTO := structural.NewAdapter(tt.args)

		assert.Equal(t, tt.want, personDTO.FullName)
	}

}
