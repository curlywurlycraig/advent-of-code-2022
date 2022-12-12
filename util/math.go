package util

func IntAbs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
