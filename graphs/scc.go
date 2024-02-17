package graphs

import (
	"github.com/samber/lo"
	"log/slog"
)

type NodeStatus[T comparable] struct {
	discover, finish int
	node             T
}

// DFSWalk persists the state of the dfs state so that we can use it computing components
type DFSWalk[T comparable] struct {
	visited map[T]bool
	// clock increases monotonically.
	clock int
	graph Graph[T]
}

func createDFSWalk[T comparable](graph Graph[T]) DFSWalk[T] {
	return DFSWalk[T]{graph: graph, clock: 0, visited: make(map[T]bool)}
}

func (d *DFSWalk[T]) time() int {
	d.clock++
	return d.clock
}

func (d *DFSWalk[T]) exploreRecursive(node T, status *[]NodeStatus[T]) {
	d.visited[node] = true
	info := NodeStatus[T]{discover: d.time(), node: node} // put discovery time
	for _, dst := range d.graph.Neighbours(node) {
		if !d.visited[dst] {
			d.exploreRecursive(dst, status)
		}
	}
	info.finish = d.time()          // done
	*status = append(*status, info) // finished nodes are added at end.
}

// Explore walks the dfs tree root at node and returns all the nodes reachable from node including itself
func (d *DFSWalk[T]) Explore(node T) []T {
	status := make([]NodeStatus[T], 0)
	d.exploreRecursive(node, &status)
	return lo.Map(status, func(item NodeStatus[T], _ int) T {
		return item.node
	})
}

// ExploreGraph traverses the entire graph starting at node.
// Returned NodeStatus contains newly finished nodes at the end of the slice, so traversing the slice in reverse order
// gives the nodes in the decreasing order of finish time
func (d *DFSWalk[T]) ExploreGraph(node T) []NodeStatus[T] {
	status := make([]NodeStatus[T], 0, d.graph.NodeCount())
	d.exploreRecursive(node, &status)
	for _, vertex := range d.graph.Nodes() {
		if !d.visited[vertex] {
			d.exploreRecursive(vertex, &status)
		}
	}
	return status
}

// ComputeSCC computes strongly connected components of a graph and returns them.
func ComputeSCC[T comparable](graph Graph[T]) [][]T {
	// step 1
	dfs := createDFSWalk(graph)
	nodeStatus := dfs.ExploreGraph(graph.Nodes()[0]) // traverse the whole graph from any node

	for _, info := range nodeStatus {
		slog.Debug("Dfs status",
			"node", info.node,
			"discover", info.discover,
			"finish", info.finish)
	}

	// step 2
	graphT := graph.Transpose()
	// traverse graph in reverse order of finish_time of original graph
	dfsT := createDFSWalk(graphT)
	allComponents := make([][]T, 0)
	componentComputed := make(map[T]bool, graph.NodeCount())

	// step 3
	for i := len(nodeStatus) - 1; i >= 0; i-- {
		node := nodeStatus[i].node
		if _, ok := componentComputed[node]; ok {
			continue
		}
		// Since we are reusing the same dfsT, already explored nodes are not traversed.
		componentNodes := dfsT.Explore(nodeStatus[i].node)
		for _, compNode := range componentNodes {
			componentComputed[compNode] = true
		}
		allComponents = append(allComponents, componentNodes)
	}

	return allComponents

}
