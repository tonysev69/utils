package utils

func Contains(a []string, b string) bool {
	for _, line := range a {
		if line == b {
			return true
		}
	}
	return false

}
