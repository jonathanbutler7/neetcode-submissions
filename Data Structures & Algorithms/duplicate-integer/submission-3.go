func hasDuplicate(nums []int) bool {
    counts := map[int]int{}
	for _, val := range nums {
		counts[val]++
		if counts[val] > 1 {
			return true
		}
	}
	return false
}
