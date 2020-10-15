package main

/*
 * 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是
 * C 的朋友。所谓的朋友圈，是指所有朋友的集合。
 *
 * 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j
 * 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。
 *
 * 示例 1:
 *
 *
 * 输入:
 * [[1,1,0],
 * ⁠[1,1,0],
 * ⁠[0,0,1]]
 * 输出: 2
 * 说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
 * 第2个学生自己在一个朋友圈。所以返回2。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * [[1,1,0],
 * ⁠[1,1,1],
 * ⁠[0,1,1]]
 * 输出: 1
 * 说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
 *
 *
 * 注意：
 *
 *
 * N 在[1,200]的范围内。
 * 对于所有学生，有M[i][i] = 1。
 * 如果有M[i][j] = 1，则有M[j][i] = 1。
 */

func union(s, w *[]int, count *int, p, q int) {
	rootP := findx(s, p)
	rootQ := findx(s, q)

	if rootP == rootQ {
		return
	}

	// 小树接到大树下面，较平衡
	if (*w)[rootP] > (*w)[rootQ] {
		(*s)[rootQ] = rootP
		(*w)[rootP] += (*w)[rootQ]
	} else {
		(*s)[rootP] = rootQ
		(*w)[rootQ] += rootP
	}

	*count--
	//fmt.Println(*count)
}

func findx(s *[]int, x int) int {
	for (*s)[x] != x {
		(*s)[x] = (*s)[(*s)[x]]
		x = (*s)[x]
	}

	return x
}

//union-find方法（并查集算法） O(logN)
func findCircleNum(M [][]int) int {
	m := len(M) //行
	if m == 0 {
		return 0
	}

	n := len(M[0]) //列

	//定义n长度的数组并且全部初始化为i
	var store []int
	var sizeWt []int
	count := m
	for i := 0; i < n; i++ {
		store = append(store, i)
		sizeWt = append(sizeWt, 1)
	}

	for i := 0; i < m; i++ {
		//内循环直到i就可以，因为是对称矩阵
		for j := i + 1; j < n; j++ {
			if M[i][j] == 0 {
				continue
			}

			if M[i][j] == 1 {
				union(&store, &sizeWt, &count, i, j)
			}
		}
	}

	return count
}

func main() {
	//x := [][]int{
	//	{1, 1, 0},
	//	{1, 1, 0},
	//	{0, 0, 1},
	//}

	x := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}
	findCircleNum(x)
}
