func twoSum(nums []int, target int) []int {
    complements := map[int]int{}
	for i, num := range nums {
		complement := target - num
		if _, exists := complements[complement]; exists {
			return []int{complements[complement], i}
		} else {
			complements[num] = i
		}
	}
	return []int{}
}
