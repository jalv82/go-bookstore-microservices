package commons

func PointerToValue[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}

	var zero T
	return zero
}

func ValueOfPointer[T any](ptr T) *T {
	return &ptr
}
