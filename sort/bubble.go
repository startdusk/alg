package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](x []T) {
	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
