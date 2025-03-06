package rotate_list

//url: https://leetcode.com/problems/rotate-array/description/

import "slices"

func rotate(nums []int, k int) []int {
	l := len(nums)
	if k > l {
		k = k % l
	}

	if k != 0 {
		fromKIndex := nums[(l - k):]
		toKIndex := slices.Clone(nums[:l-k])

		for i := 0; i < k; i++ {
			nums[i] = fromKIndex[i]
		}

		for i := 0; i < l-k; i++ {
			nums[k+i] = toKIndex[i]

		}
	}

	return nums
}
