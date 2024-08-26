package dylan_algo

// https://leetcode.cn/problems/two-sum/
// [2,7,11,15] 9
func twoSum(num []int, target int) []int {
	cacheMap := map[int]int{}
	for i, curNum := range num {
		subNum := target - curNum
		subNumIndex, ok := cacheMap[subNum]
		if ok {
			return []int{subNumIndex, i}
		}
		cacheMap[curNum] = i
	}
	return nil

}
