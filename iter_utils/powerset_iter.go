package iter_utils

import (
	"iter"
	"math"
)

func PowerSetIter[T any](values []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < int(math.Pow(2, float64(len(values)))); i++ {
			result := make([]T, 0)
			for j := 0; j < len(values); j++ {
				if i&(1<<j) != 0 {
					result = append(result, values[j])
				}
			}
			if !yield(result) {
				return
			}
		}
	}
}

func PowerSet[T any](values []T) [][]T {
	if len(values) == 0 {
		return nil
	}
	results := make([][]T, 0)
	for i := 0; i < int(math.Pow(2, float64(len(values)))); i++ {
		result := make([]T, 0)
		for j := 0; j < len(values); j++ {
			if i&(1<<j) != 0 {
				result = append(result, values[j])
			}
		}
		results = append(results, result)
	}
	return results
}
