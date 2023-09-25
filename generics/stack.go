package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.values) - 1
	result := s.values[index]
	s.values = s.values[:index]
	return result, true
}
