import "slices"
func groupAnagrams(words []string) [][]string {
	strMap := map[string][]string{}
	for _, word := range words {
		runes := []rune(word)
		slices.Sort(runes)
		sortedStr := string(runes)
		if _, exists := strMap[sortedStr]; exists {
			strMap[sortedStr] = append(strMap[sortedStr], word)
		} else {
			strMap[sortedStr] = []string{word}
		}
	}
	result := [][]string{}
	for _, val := range strMap {
		result = append(result, val)
	}
	return result
}
