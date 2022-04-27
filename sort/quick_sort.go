package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](x []T) {
	quickSort(x, 0, len(x)-1)
}

func quickSort[T constraints.Ordered](x []T, low, high int) []T {
	if low < high {
		var p int
		x, p = partition(x, low, high)
		x = quickSort(x, low, p-1)
		x = quickSort(x, p+1, high)
	}
	return x
}

func partition[T constraints.Ordered](x []T, low, high int) ([]T, int) {
	pivot := x[high]
	i := low
	for j := low; j < high; j++ {
		if x[j] < pivot {
			x[i], x[j] = x[j], x[i]
			i++
		}
	}
	x[i], x[high] = x[high], x[i]
	return x, i
}
