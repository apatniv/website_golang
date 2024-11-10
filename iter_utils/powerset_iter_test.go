package iter_utils_test

import (
	"github.com/apatniv/website_golang/iter_utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
	"testing"
)

func TestPowerSetIter(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{
			name:  "Empty values",
			input: []int{},
			want:  [][]int{{}},
		},
		{
			name:  "1 value",
			input: []int{100},
			want:  [][]int{{}, {100}},
		},
		{
			name:  "2 values",
			input: []int{100, 99},
			want:  [][]int{{}, {100}, {99}, {99, 100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for result := range iter_utils.PowerSetIter(tt.input) {
				slices.Sort(result)
				require.Contains(t, tt.want, result)
			}
		})
		t.Run("iterative version"+tt.name, func(t *testing.T) {
			for _, result := range iter_utils.PowerSet(tt.input) {
				slices.Sort(result)
				require.Contains(t, tt.want, result)
			}
		})
	}
}
