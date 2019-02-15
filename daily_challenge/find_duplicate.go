package daily_challenge

import "fmt"

/**
 * daily-challenge-287
 * PUZZLE: Find the Duplicate Number
 */
func findDuplicate(nums []int) int {
	target := 0
	for i, val := range nums {
		isFind := false
		j := i + 1
		for ; j < len(nums); j++ {
			if val == nums[j] {
				isFind = true
				break
			}
		}

		if isFind {
			target = val
			break
		}
	}
	return target
}

/**
 * 通过二分法减少检索的次数，计算小于某个值A的数字B的多少C；
 * 如果C不大于A，那么一定是大于A的数字有重复；反之就是小于A的数字重复；
 */
func findDuplicateOptimize(nums []int) int {
	// target := 0
	fmt.Printf("%v\n", nums)

	low := 1
	high := len(nums)
	for low < high {
		middle := low + (high-low)/2

		count := 0
		for i := range nums {
			if nums[i] <= middle {
				count++
			}
		}

		println(low, high, middle, count)
		if count <= middle {
			low = middle + 1
		} else {
			high = middle
		}
	}
	return low
}
