package iter_utils

import "iter"

func SliceToIter[V any](values []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}
