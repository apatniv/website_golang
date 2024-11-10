package iter_utils

import (
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func TestLimit(t *testing.T) {
	input := []int{1, 2, 3}
	tests := []struct {
		name  string
		limit int
		want  []int
	}{
		{
			name:  "0 limit",
			limit: 0,
			want:  nil,
		},
		{
			name:  "1 limit",
			limit: 1,
			want:  []int{1},
		},
		{
			name:  "negative limit",
			limit: -10,
			want:  nil,
		},
		{
			name:  "More than the size of sequence",
			limit: 100,
			want:  []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(Limit(tt.limit, SliceToIter(input)))
			require.Equal(t, tt.want, got)
		})
	}
}
