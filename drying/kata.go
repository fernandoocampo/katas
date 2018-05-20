package drying

func Potatoes(p0, w0, p1 int) int {
	/*
		inverse proportionality
		p0 -> w0
		p1 -> x
		x = w0 . p0 / p1

		w0 * (100 - p0) / (100 - p1)
		100 * ((100) - 99) / ((100) - 98)
		127 * (100 - 82) / (100 - 80)
	*/
	return int(w0 * (100 - p0) / (100 - p1))
}
