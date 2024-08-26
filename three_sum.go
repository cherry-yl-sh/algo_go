package dylan_algo

import (
	"sort"
)

// https://leetcode.cn/problems/3sum/
// [-1,0,1,2,-1,-4]
func threeSum(num []int) [][]int {
	sort.Ints(num)
	res := make([][]int, 0)

	for k := 0; k < len(num)-1; k++ {
		if num[k] > 0 {
			break
		}
		if k > 0 && num[k] == num[k-1] {
			continue
		}
		i := k + 1
		j := len(num) - 1
		for i < j {
			sum := num[k] + num[i] + num[j]
			if sum < 0 {
				for i < j && num[i] == num[i+1] {
					i++
				}
				i++
			} else if sum > 0 {
				for i < j && num[j] == num[j-1] {
					j--
				}
				j--
				continue
			} else {
				res = append(res, []int{num[k], num[i], num[j]})
				i++
				j--
				for i < j && num[i] == num[i+1] {
					i++
				}
				for i < j && num[j] == num[j-1] {
					j--
				}
			}
		}

	}
	return res
}
