package util

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func OrderPoints(p1, p2 int) (int, int) {
	if p1 <= p2 {
		return p1, p2
	}
	return p2, p1
}
