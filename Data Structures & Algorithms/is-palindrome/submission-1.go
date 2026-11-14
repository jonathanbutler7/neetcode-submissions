func toAlphanumericASCII(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func isPalindrome(s string) bool {
	result := toAlphanumericASCII(strings.ToLower(strings.ReplaceAll(s, " ", "")))
	left := 0
	right := len(result) - 1
	fmt.Println(right, left)
	for left < right {

		fmt.Println("left", string((result[left])))
		fmt.Println("right", string((result[right])))
		if result[left] != result[right] {
			return false
		}
		left++
		right--
	}
	return true	
}
