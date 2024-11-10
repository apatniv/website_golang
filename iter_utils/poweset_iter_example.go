package iter_utils

import "fmt"

func PowerSetExample() {
	input := []int{9, 50, 99}
	fmt.Printf("Computing powerset for=%v\n", input)
	for value := range PowerSetIter([]int{9, 50, 99}) {
		fmt.Printf("value = %v\n", value)
	}
}
