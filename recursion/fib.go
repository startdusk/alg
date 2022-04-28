package recursion

// 递归实现Fibonacci序列
func Fib(n int) int {
	if n <= 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// 循环实现Fibonacci序列
func FibWithLoop(n int) int {
	var a, b = 0, 1
	for i := 0; i < n; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return b
}

// 动态规划实现Fibonacci序列
func FibWithDynamic(n int) int {
	res, _ := fibWithDynamic(n)
	return res
}

func fibWithDynamic(n int) (int, int) {
	if n == 0 {
		return 1, 0
	}

	a, b := fibWithDynamic(n - 1)
	return a + b, a
}

// 斐波那契数列使用尾递归实现
// 如果一个函数返回自身递归调用的结果，那么调用过程会被替换为一个循环，它可以显著提高速度。
// 尾递归是一种在函数的最后执行递归调用语句的特殊形式的递归
// 尾递归就是从最后开始计算, 每递归一次就算出相应的结果, 也就是说, 函数调用出现在调用者函数的尾部, 因为是尾部, 所以根本没有必要去保存任何局部变量。
// 直接让被调用的函数返回时越过调用者,返回到调用者的调用者去。
// 初始值 prev = 1, res = 1
func FibWithTailRecursice(n int) int {
	return fibWithTailRecursice(n, 1, 1)
}

func fibWithTailRecursice(n, prev, res int) int {
	if n == 0 {
		return prev
	}

	return fibWithTailRecursice(n-1, res, prev+res)
}
