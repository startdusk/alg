package sort

import "golang.org/x/exp/constraints"

// 归并排序的时间复杂度任何情况下都是 O(nlogn)，看起来非常优秀。
// （待会儿你会发现，即便是快速排序，最坏情况下，时间复杂度也是 O(n2)。）
// 但是，归并排序并没有像快排那样，应用广泛，这是为什么呢？因为它有一个致命的“弱点”，那就是归并排序不是原地排序算法。
func MergeSort[T constraints.Ordered](x []T) []T {
	if len(x) < 2 {
		return x
	}
	mid := (len(x)) / 2
	return merge(MergeSort(x[:mid]), MergeSort(x[mid:]))
}

// 每次归并都会生成一个新的数组
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
