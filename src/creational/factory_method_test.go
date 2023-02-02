package creational_test

import (
	"testing"

	"github.com/DioCoding/desing-patterns/src/creational"
	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {

	type Arg struct {
		typeGun string
	}

	tests := []struct {
		name    string
		arg     Arg
		want    string
		wantErr string
	}{
		{
			name: "factory revolver",
			arg: Arg{
				typeGun: "revolver",
			},
			want: "revolver",
		},
		{
			name: "factory ak47",
			arg: Arg{
				typeGun: "AK47",
			},
			want: "Ak47",
		},
		{
			name: "factory invalid",
			arg: Arg{
				typeGun: "invalid",
			},
			wantErr: "invalid type gun",
		},
	}

	for _, tt := range tests {
		gun, err := creational.NewFactoryMethod(tt.arg.typeGun)

		if len(tt.wantErr) > 0 {
			assert.Equal(t, tt.wantErr, err.Error())
		} else {
			assert.Equal(t, tt.want, gun.Name())
		}
	}

}
