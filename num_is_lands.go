package dylan_algo

// https://leetcode.cn/problems/number-of-islands/description/
var (
	land byte = '1'
)

func numIslands(grid [][]byte) int {
	landCount := 0
	if len(grid) == 0 {
		return landCount
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				landCount++
				dfs(grid, i, j)
			}
		}
	}
	return landCount
}
func dfs(grid [][]byte, x int, y int) {

	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return
	}
	if grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
