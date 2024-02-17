package graphs

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
	"testing"
)

func sortComponents(components [][]string) {
	for _, component := range components {
		slices.Sort(component)
	}
}

func TestComputeSCC(t *testing.T) {
	tests := []struct {
		name            string
		rawEdges        [][2]string
		additionalNodes []string
		want            [][]string
	}{
		{
			name: "linear_chain",
			rawEdges: [][2]string{
				{"a", "b"},
				{"b", "c"},
			},
			want: [][]string{
				{"a"},
				{"b"},
				{"c"},
			},
		},
		{
			name:            "single isolated nodes",
			additionalNodes: []string{"a", "b"},
			want: [][]string{
				{"a"},
				{"b"},
			},
		},
		{
			name: "complete 3 node cycle",
			rawEdges: [][2]string{
				{"a", "b"},
				{"b", "c"},
				{"c", "a"},
			},
			want: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "4 components example",
			rawEdges: [][2]string{
				{"a", "b"},
				{"b", "a"},
				{"c", "d"},
				{"d", "c"},
				{"a", "c"},
				{"b", "d"},
				{"d", "e"},
				{"d", "f"},
			},
			want: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e"},
				{"f"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edges := lo.Map(tt.rawEdges, func(rawEdge [2]string, _ int) Edge[string] {
				return Edge[string]{Src: rawEdge[0], Dst: rawEdge[1]}
			})
			graph := CreateGraph(edges, tt.additionalNodes...)
			got := ComputeSCC(graph)
			require.Equal(t, len(tt.want), len(got))
			sortComponents(got)
			sortComponents(tt.want)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
