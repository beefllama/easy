package easy

func GetPointer[T any](ptr *T) (val T) {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func PointerTo[T any](val T) *T {
	return &val
}
