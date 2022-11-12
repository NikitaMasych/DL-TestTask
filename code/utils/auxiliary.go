package utils

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsSlice[T comparable](s [][]T, e []T) bool {
	for _, v := range s {
		if isEqual(v, e) {
			return true
		}
	}
	return false
}

func isEqual[T comparable](a, b []T) bool {
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
