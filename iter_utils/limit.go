package iter_utils

import (
	"iter"
)

func Limit[V any](limit int, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		if limit <= 0 {
			return
		}
		count := 0
		for v := range seq {
			count++
			if !yield(v) || count == limit {
				return
			}
		}
	}
}
