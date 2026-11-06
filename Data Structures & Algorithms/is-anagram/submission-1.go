func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := map[rune]int{}
	for _, letter := range s {
		counts[letter]++
	}
	for _, letter := range t {
		counts[letter]--
		if counts[letter] < 0 {
			return false
		}
	}
	return true
}
