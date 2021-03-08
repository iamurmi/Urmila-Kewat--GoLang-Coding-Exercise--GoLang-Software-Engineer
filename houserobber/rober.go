package houserobber

// Rob ...
func Rob(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}

		return nums[1]
	}

	cache1 := make(map[int]int)
	cache2 := make(map[int]int)

	return robMax(robRecursive(nums[:n-1], n-2, cache1), robRecursive(nums[1:], n-2, cache2))
}

func robRecursive(nums []int, index int, cache map[int]int) int {
	// memory checking
	if val, found := cache[index]; found {
		return val
	}
	// BASE CONDITION
	if index == 0 {
		cache[index] = nums[0]
		return cache[index]
	}

	if index == 1 {
		max := robMax(nums[0], nums[1])
		cache[index] = max
		return cache[index]
	}

	// XXXXXXX BASE CONDITION XXXXXXXX

	// DP recursion
	max := robMax(robRecursive(nums, index-1, cache), robRecursive(nums, index-2, cache)+nums[index])
	cache[index] = max

	return cache[index]
}

func robMax(sum1 int, sum2 int) int {
	if sum2 > sum1 {
		return sum2
	}

	return sum1
}
