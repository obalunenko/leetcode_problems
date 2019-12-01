package twosum

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length == 1 {
		if nums[0] == target {
			return []int{0}
		}
	}

	for i := 0; i < length; i++ {
		a := nums[i]
		for j := i + 1; j < length; j++ {
			b := nums[j]

			sum := a + b

			if sum == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
