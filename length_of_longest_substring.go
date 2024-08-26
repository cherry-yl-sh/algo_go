package dylan_algo

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	res := 0
	set := map[byte]int{}

	for left, right := 0, 0; right < len(s); right++ {
		c := s[right]
		// left move
		_, ok := set[s[right]]
		for ok == true {
			delete(set, s[right])
			_, ok = set[s[right]]
			left++
		}

		res = max(res, right-left+1)
		set[c] = right
	}
	return res
}
