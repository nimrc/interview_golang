package main

import "fmt"

/*
 * Q: 找出数组中的重复数字
 *
 * T = O(n)
 * 边界：数组为空
 */

func duplicate(nums []int) (bool, int) {
	if len(nums) == 0 {
		return false, 0
	}

	var dup int

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				dup = nums[i]
				return true, dup
			}

			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return false, 0
}

/*
 * Q: 不修改元素的解法
 *
 * 思路: 二分查找
 * T = O(nlogn)
 */
func duplicateUnchange(nums []int) (bool, int) {
	if len(nums) == 0 {
		return false, 0
	}

	end := len(nums) - 1
	start := 1

	for end >= start {
		mid := ((end - start) / 2) + start

		counter := 0
		for _, i := range nums {
			if i <= mid && i >= start {
				counter++
			}
		}

		if end == start {
			if counter > 1 {
				return true, start
			} else {
				break
			}
		}

		if counter > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}

	}

	return false, 0
}

func main() {
	nums := []int{2, 3, 1, 0, 4, 2}
	res, dup := duplicate(nums)

	fmt.Println(res, dup, nums)

	nums = []int{2, 3, 5, 4, 2, 1, 6, 7}

	res, dup = duplicateUnchange(nums)

	fmt.Println(res, dup)
}
