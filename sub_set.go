package dylan_algo

// https://leetcode.cn/problems/subsets/description/
func subsets(num []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	var dfs func(int)
	dfs = func(level int) {
		if level == len(num) {
			tempTemp := make([]int, 0)
			tempTemp = append(tempTemp, temp...)
			res = append(res, tempTemp)
			return
		}
		temp = append(temp, num[level])
		dfs(level + 1)
		temp = temp[:len(temp)-1]
		dfs(level + 1)

	}
	dfs(0)
	return res
}
