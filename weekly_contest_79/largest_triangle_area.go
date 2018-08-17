package weekly_contest_79

import "math"

func isSamePoint(point1 []int, point2 []int) bool {
	if point1[0] == point2[0] && point1[1] == point2[1] {
		return true
	}
	return false
}

func isTriangle(point1 []int, point2 []int, point3 []int) bool {
	if point1[0] == point2[0] && point3[0] == point2[0] {
		return false
	} else if point1[1] == point2[1] && point3[1] == point2[1] {
		return false
	} else if isSamePoint(point1, point2) || isSamePoint(point1, point3) || isSamePoint(point2, point3) {
		return false
	}

	return true
}

func lineLength(point1 []int, point2 []int) float64 {
	height := point1[1] - point2[1]
	width := point1[0] - point2[0]

	area := float64(height*height + width*width)
	return math.Sqrt(area)
}

func triangleArea(a float64, b float64, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				if !isTriangle(points[i], points[j], points[k]) {
					continue
				}

				a := lineLength(points[i], points[j])
				b := lineLength(points[j], points[k])
				c := lineLength(points[i], points[k])

				current := triangleArea(a, b, c)
				if current > maxArea {
					maxArea = current
				}
			}
		}
	}
	return maxArea
}
