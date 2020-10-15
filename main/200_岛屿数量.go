package main

import "fmt"

func numIslands(grid [][]byte) int {
	/*
		以陆地(i,j)为中心扩散，进行深度优先遍历，找到与其连通的点，将其都标记为0，视为一个岛屿。新的岛屿是以新的陆地(i,j)开始，找到与其连通的点构建为一个岛屿。
		o(m*n)
	*/
	count := 0
	for line := range grid {
		for col := range grid[line] {
			if grid[line][col] == '0' || grid[line][col] == '2' {
				continue
			}

			count++
			dfs33(grid, line, col)
		}
	}
	return count
}

func dfs33(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs33(grid, i-1, j)
	dfs33(grid, i+1, j)
	dfs33(grid, i, j-1)
	dfs33(grid, i, j+1)
}

func main(){
	grid := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'}}

	fmt.Print(numIslands(grid))
}
