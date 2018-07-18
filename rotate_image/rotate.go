package rotate_image

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		innerLen := len(matrix) - i - 1
		distance := len(matrix) - i - 1

		for j := i; j < innerLen; j++ {
			println(i, j, distance)
			if j == distance && i == distance {
				continue
			}
			println("++", matrix[i][j], matrix[distance-j][i], matrix[distance-i][distance-j], matrix[j][distance-i])
			temp := matrix[i][j]
			matrix[i][j] = matrix[distance-j][i]
			matrix[distance-j][i] = matrix[distance-i][distance-j]
			matrix[distance-i][distance-j] = matrix[j][distance-i]
			matrix[j][distance-i] = temp
		}
	}
}
