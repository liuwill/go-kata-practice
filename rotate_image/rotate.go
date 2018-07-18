package rotate_image

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		distance := len(matrix) - i - 1

		for j := i; j < distance; j++ {

			temp := matrix[i][j]
			matrix[i][j] = matrix[distance-j+i][i]
			matrix[distance-j+i][i] = matrix[distance][distance-j+i]
			matrix[distance][distance-j+i] = matrix[j][distance]
			matrix[j][distance] = temp
		}
	}
}
