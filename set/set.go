package set

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(element T) {
	(*s)[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(*s, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := (*s)[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) ToSlice() []T {
	list := make([]T, 0, len(*s))
	for key := range *s {
		list = append(list, key)
	}
	return list
}

func NewSetFromSlice[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	for _, val := range slice {
		set.Add(val)
	}
	return set
}

func (s *Set[T]) Copy() Set[T] {
	newSet := make(Set[T])
	for key, val := range *s {
		newSet[key] = val
	}
	return newSet
}

func (s *Set[T]) IsDisjoint(otherSet Set[T]) bool {
	for attr := range *s {
		if otherSet.Contains(attr) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Intersect(otherSet Set[T]) Set[T] {
	intersect := make(Set[T])
	for val := range *s {
		if otherSet.Contains(val) {
			intersect.Add(val)
		}
	}
	return intersect
}
