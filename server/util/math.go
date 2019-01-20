package util

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func OrderPoints(px1, py1, px2, py2 int) (rx1, ry1, rx2, ry2 int) {
	if px1 <= px2 {
		rx1 = px1
		ry1 = py1
		rx2 = px2
		ry2 = py2
	} else {
		rx1 = px1
		ry1 = py1
		rx2 = px2
		ry2 = py2
	}
	return
}
