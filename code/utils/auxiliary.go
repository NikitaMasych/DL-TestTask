package utils

func Contains(list [][]string, element []string) bool {
	for _, current := range list {
		if isEqual(current, element) {
			return true
		}
	}
	return false
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
