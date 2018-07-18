package rotate_image

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		innerLen := len(matrix) - i - 1
		for j := i; j < innerLen; j++ {
			distance := len(matrix) - j - 1
			if i == distance {
				continue
			}
			temp := matrix[i][j]
			matrix[i][j] = matrix[distance][j]
			matrix[distance][j] = matrix[distance][distance]
			matrix[distance][distance] = matrix[i][distance]
			matrix[i][distance] = temp
		}
	}
}
