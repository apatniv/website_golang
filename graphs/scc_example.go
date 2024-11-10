package graphs

import "log/slog"

// SccExample runs a strongly connected algorithm on a simple graph
func SccExample() {
	slog.Info("Running the SCC example")
	edges := []Edge[string]{
		{"a", "b"},
		{"b", "a"},
		{"c", "d"},
		{"d", "c"},
		{"a", "c"},
		{"b", "d"},
	}
	graph := CreateGraph(edges)
	components := ComputeSCC(graph)
	slog.Info("SCC Results", "components", components)
}
