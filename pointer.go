package easy

// GetPointer returns a value ptr points to, or zero value if ptr is nil.
func GetPointer[T any](ptr *T) (val T) {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// PointerTo returns a pointer to the value.
func PointerTo[T any](val T) *T {
	return &val
}
