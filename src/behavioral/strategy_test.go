package behavioral_test

import (
	"testing"

	"github.com/DioCoding/desing-patterns/behavioral/strategies"
	"github.com/stretchr/testify/assert"
)

func TestStrategiesWithStruct(t *testing.T) {
	type Args struct {
		strategy strategies.SaveFile
		text     string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "local estrategy",
			args: Args{
				strategy: strategies.SaveLocal{},
				text:     "Test_file",
			},
			want: "Save local: Test_file",
		},
		{
			name: "cloud estrategy",
			args: Args{
				strategy: strategies.SaveCloud{},
				text:     "Test_file",
			},
			want: "Save cloud: Test_file",
		},
	}
	for _, tt := range tests {
		strategy := strategies.New(tt.args.text, tt.args.strategy)
		assert.Equal(t, tt.want, strategy.Print())
	}
}

func TestStrategiesWithFunc(t *testing.T) {
	type Args struct {
		strategyName string
		text         string
	}

	tests := []struct {
		name    string
		args    Args
		want    string
		wantErr string
	}{
		{
			name: "local estrategy",
			args: Args{
				strategyName: "local",
				text:         "Test_file",
			},
			want: "Save local: Test_file",
		},
		{
			name: "cloud estrategy",
			args: Args{
				strategyName: "cloud",
				text:         "Test_file",
			},
			want: "Save cloud: Test_file",
		},
		{
			name: "invalid estrategy",
			args: Args{
				strategyName: "xsxsaxasx",
				text:         "Test_file",
			},
			wantErr: "invalid type",
		},
	}
	for _, tt := range tests {
		strategy, err := strategies.Strategy(tt.args.strategyName)
		if tt.wantErr != "" {
			assert.Equal(t, err.Error(), tt.wantErr)
		} else {
			assert.Equal(t, tt.want, strategy(tt.args.text))
		}
	}

}

// How to skip test.
func TestStrategySkip(t *testing.T) {
	t.Skip("example to skip test.")

	assert.Equal(t, 50, 50)
}
