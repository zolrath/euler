/*
Problem 18: Maximum Path Sum I
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
3
7 4
2 4 6
8 5 9 3
That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle [py]
*/

package problems

const ANSWER_018 = 1074

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reduceRow(row []int) []int {
	newrow := []int{}
	for i := 0; i < len(row)-1; i++ {
		newrow = append(newrow, max(row[i], row[i+1]))
	}
	return newrow

}
func maxPath(tri [][]int) int {
	for l := len(tri) - 1; l > 0; l-- {
		row := reduceRow(tri[l])
		rowAbove := tri[l-1]
		for i, _ := range rowAbove {
			rowAbove[i] += row[i]
		}
	}
	return tri[0][0]
}

func Euler018() int {
	py := [][]int{
		[]int{75},
		[]int{95, 64},
		[]int{17, 47, 82},
		[]int{18, 35, 87, 10},
		[]int{20, 4, 82, 47, 65},
		[]int{19, 1, 23, 75, 3, 34},
		[]int{88, 2, 77, 73, 7, 63, 67},
		[]int{99, 65, 4, 28, 6, 16, 70, 92},
		[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
		[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		[]int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}
	return maxPath(py)
}
