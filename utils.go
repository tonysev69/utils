package utils

func InSlice(a []string, b string) bool {
	for _, line := range a {
		if line == b {
			return true
		}
	}
	return false

}
func InSliceInt(a []int, b int) bool {
	for _, line := range a {
		if line == b {
			return true
		}
	}
	return false
}
