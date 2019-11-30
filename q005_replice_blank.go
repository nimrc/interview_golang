package main

import (
	"fmt"
	"strings"
)

/*
 * 替换字符串中的空格
 *
 * 思路 从后往前遍历
 * T = O(n)
 */
func replaceBlank(str string) string {
	counter := 0
	for _, c := range str {
		if c == ' ' {
			counter++
		}
	}

	str += strings.Repeat(" ", counter*2)
	runes := []rune(str)
	i := len(runes) - 1
	j := i - counter*2

	for i >= 0 {
		if runes[j] != ' ' {
			runes[i] = runes[j]
			i--
			j--
		} else {
			runes[i] = '0'
			runes[i-1] = '2'
			runes[i-2] = '%'
			i -= 3
			j--
		}
	}

	return string(runes)
}

func main() {
	str := "We are happy."

	str = replaceBlank(str)

	fmt.Println(str)
}
