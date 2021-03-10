package main

import "fmt"

/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*
	设定初始条件。第一行中obstacleGrid第一个1之前的每个点的路径和均为1，此点之后为0；第一列同理
	开始动态规划。遍历obstacleGrid中的每一个点。对于obstacleGrid[i][j]==1的点，说明不能到此点，
	dp[i][j]=0;对于obstacleGrid[i][j]==0的点，说明可以到达这一点，则使用递推公式。
	返回dp[row-1][col-1]
	*/
	if len(obstacleGrid) == 0 {
		return 0
	}
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	//1 设定初始条件
	i := 0
	for i < row && obstacleGrid[i][0] == 0 {
		dp[i][0] = 1
		i++
	}
	if i != row {
		for j := i; j < row; j++ {
			dp[j][0] = 0
		}
	}

	i = 0
	for i < col && obstacleGrid[0][i] == 0 {
		dp[0][i] = 1
		i++
	}
	if i != col {
		for j := i; j < col; j++ {
			dp[0][j] = 0
		}
	}

	//2 开启动态规划
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 { //对于不能走的点，dp为0
				dp[i][j] = 0
			} else { //对于可以走的点，dp用地推公式求解
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[row-1][col-1]
}

func main() {
	x := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(x))
}
