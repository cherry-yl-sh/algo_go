package dylan_algo

// https://leetcode.cn/problems/jump-game/description/
func canJump(nums []int) bool {
	rightMax := 0

	for i, num := range nums {
		if i <= rightMax {
			reach := i + num
			rightMax = max(rightMax, reach)
			if rightMax >= len(nums)-1 {
				return true
			}
		}

	}
	return false
}
