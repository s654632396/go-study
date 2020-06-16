package main


func main() {

	// 给定一个 n × n 的二维矩阵表示一个图像
	// 将图像顺时针旋转 90 度

	n := 5
	m := make([][]int, n)
	var data = 1
	for i, _ := range m {
		m[i] = make([]int, n)
		for x := 0; x < n; x++ {
			m[i][x] = data
			data++
		}
	}
	rotate(m)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i, _ := range matrix {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i, _ := range matrix {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

