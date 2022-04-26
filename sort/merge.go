package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](x []T) []T {
	if len(x) < 2 {
		return x
	}
	mid := (len(x)) / 2
	return merge(MergeSort(x[:mid]), MergeSort(x[mid:]))
}

func merge[T constraints.Ordered](left, right []T) []T {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]T, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
