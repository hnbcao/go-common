package _29

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// calc matrix width and height
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	total := len(matrix) * len(matrix[0])
	res := make([]int, total)
	x, y := 0, 0
	signD := r > 0 // xy，true 计算y
	signR := true  // 加减
	for i := 0; i < total; i++ {
		res[i] = matrix[x][y]
		// 顺序Y轴
		if signD && signR {
			y++
			if y == r {
				t++
				signD = !signD
			}
		} else
		// 顺序X轴
		if !signD && signR {
			x++

			if x == b {
				r--
				signR = !signR
				signD = !signD
			}
		} else
		// 倒序Y轴
		if signD && !signR {
			y--
			if y == l {
				b--
				signD = !signD
			}
		} else
		// 倒序X轴
		if !signD && !signR {
			x--
			if x == t {
				l++
				signR = !signR
				signD = !signD
			}
		}
	}

	return res
}
