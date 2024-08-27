package dylan_algo

func searchMatrix(matrix [][]int, target int) bool {
	rowNum := findRow(matrix, target)
	if rowNum < 0 {
		return false
	}
	index := searchInRow(matrix[rowNum], target)
	if index < 0 {
		return false
	}
	return true
}
func findRow(matrix [][]int, target int) int {
	left, right := 0, len(matrix)-1
	if matrix[0][0] > target {
		return -1
	}
	mlen := len(matrix)
	if target > matrix[mlen-1][0] {
		return mlen - 1
	}
	mid := (right-left)/2 + left
	for left < right {
		if matrix[left][0] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
func searchInRow(row []int, target int) int {
	left, right := 0, len(row)-1
	for left < right {
		mid := (right-left)/2 + left
		if row[mid] == target {
			return mid
		} else if row[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
