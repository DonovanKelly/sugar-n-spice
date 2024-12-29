package all

func All[T comparable](slice []T, condition func(T) bool) bool {
	for _, v := range slice {
		if !condition(v) {
			return false
		}
	}
	return true
}
