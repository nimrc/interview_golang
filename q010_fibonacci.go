package main

import "fmt"

/*
 * Q: 斐波那契数列 递归/动态规划写法
 */
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciDp(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	x, y := 1, 1

	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("fibonacci(%d) = %d; fibonacciDp(%d) = %d\n", i, fibonacci(i), i, fibonacciDp(i))
	}

}
