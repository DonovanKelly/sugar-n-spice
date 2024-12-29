package any

func Any[T comparable](slice []T, condition func(T) bool) bool {
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}
