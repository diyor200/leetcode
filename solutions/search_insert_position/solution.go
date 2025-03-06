package searchinsertposition

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		}

		return 0
	} else if len(nums) == 0 {
		return 0
	}

	var index = 0
	for {
		middle := len(nums) / 2
		if nums[middle] == target {
			return index + middle
		} else if nums[middle] < target {
			nums = nums[middle:]
			index += middle
		} else {
			nums = nums[:middle]
		}

		if len(nums) == 1 {
			if nums[0] < target {
				return index + 1
			}

			return index
		}
	}
}
