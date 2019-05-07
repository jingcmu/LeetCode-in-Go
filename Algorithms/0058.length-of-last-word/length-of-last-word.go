package problem0058

func lengthOfLastWord(s string) int {
    strs := strings.Split(strings.TrimSpace(s), " ")
    return len(strs[len(strs) - 1])
}
//
func lengthOfLastWord(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	start := size - 1
	for ; start >= 0 && s[start] == ' '; start-- {
	}

	res := 0
	for i := start; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		res++
	}

	return res
}
