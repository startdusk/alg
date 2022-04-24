package sort

import "golang.org/x/exp/constraints"

// 插入排序，x 表示数组
func InsertSort[T constraints.Ordered](x []T) {
	n := len(x)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		v := x[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if x[j] > v {
				x[j+1] = x[j] // 数据移动，这步就是为什么插入排序比冒泡排序更受欢迎，插入这里只有一步操作，而冒泡有三步，会耗费世间，虽然它两的复杂度都一样
			} else {
				break
			}
		}
		x[j+1] = v // 插入数据
	}
}
