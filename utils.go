package go_cc

import "strconv"

func Float64(str string) (num float64) {
	num, _ = strconv.ParseFloat(str, 64)
	return
}

func Int(str string) (num int) {
	num, _ = strconv.Atoi(str)
	return
}

func hcf(a, b int) int {
	if b == 0 {
		return a
	} else {
		return hcf(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / hcf(a, b)
}
