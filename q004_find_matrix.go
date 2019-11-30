package main

import "fmt"

/*
 * 二维数组中查找元素
 *
 */

func find(matrix [][]int, num int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		if matrix[row][col] == num {
			return true
		} else if matrix[row][col] > num {
			col--
		} else {
			row++
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	num := 7

	fmt.Println(find(matrix, num))
}
