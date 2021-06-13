package algebra

// Extended Euclidean Algorithm

func EEA(a, b int64) (int64, int64, int64) {
	var x, y, x1, y1 int64 = 1, 0, 0, 1
	for b != 0 {
		q := a / b
		x, x1 = x1, x-q*x1
		y, y1 = y1, y-q*y1
		a, b = b, a-q*b
	}
	return a, x, y
}
