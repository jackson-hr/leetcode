package medium

func ShortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	// 起点终点不可达
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	// 初始化
	dest := make([][]int, n)
	for i := 0; i < n; i++ {
		dest[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dest[i][j] = 0xffffffff
		}
	}
	dest[0][0] = 1
	// 八个方向可达坐标, 起点开始
	geo := [][2]int{{0, 0}}
	x, y := 0, 0
	// 处理目标不可达
	for len(geo) > 0 {
		x, y = geo[0][0], geo[0][1]
		// 移除已经走过的坐标
		geo = geo[1:]
		// 到达目的地
		if x == n-1 && y == n-1 {
			return dest[x][y]
		}
		// 控制八个方向
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				// 越界处理
				if x+i < 0 || y+j < 0 || x+i >= n || y+j >= n {
					continue
				}
				// 目标不可达或距离更远
				if grid[x+i][y+j] == 1 || dest[x][y]+1 >= dest[x+i][y+j] {
					continue
				}
				dest[x+i][y+j] = dest[x][y] + 1
				geo = append(geo, [2]int{x + i, y + j})
			}
		}
	}
	return -1
}
