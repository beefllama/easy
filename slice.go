package easy

// Map converts a slice of one type to a slice of another type.
func Map[Tsrc any, Tdest any](s []Tsrc, convertFunc func(e Tsrc) Tdest) []Tdest {
	converted := make([]Tdest, 0, len(s))

	for _, element := range s {
		converted = append(converted, convertFunc(element))
	}

	return converted
}

// Filter returns a slice of elements for which filterFunc returns true.
func Filter[T any](s []T, filterFunc func(e T) bool) []T {
	filtered := make([]T, 0, len(s))

	for _, element := range s {
		if filterFunc(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}

// MapFilter filters a slice and maps its elements to another type (original slice is unaffected).
func MapFilter[Tsrc any, Tdest any](s []Tsrc, convertAndFilterFunc func(e Tsrc) (Tdest, bool)) []Tdest {
	converted := make([]Tdest, 0, len(s))

	for _, element := range s {
		convertedElement, ok := convertAndFilterFunc(element)
		if ok {
			converted = append(converted, convertedElement)
		}
	}

	return converted
}

// Contains returns true if s contains element.
func Contains[T comparable](s []T, element T) bool {
	for _, e := range s {
		if e == element {
			return true
		}
	}
	return false
}

// ContainsFunc returns true if s has element for which testFunc returns true.
func ContainsFunc[T comparable](s []T, testFunc func(e T) bool) bool {
	for _, e := range s {
		if testFunc(e) {
			return true
		}
	}
	return false
}

// Batch splits up a slice into batches of batchSize length.
func Batch[T any](s []T, batchSize int) [][]T {
	if len(s) == 0 {
		return [][]T{}
	}

	if len(s) <= batchSize || batchSize <= 0 {
		return [][]T{s}
	}

	batchesAmount := Ternary(len(s)%batchSize == 0, len(s)/batchSize, len(s)/batchSize+1)
	batches := make([][]T, 0, batchesAmount)

	for i := range batchesAmount {
		if i == batchesAmount-1 {
			batches = append(batches, s)
		} else {
			batches = append(batches, s[:batchSize])
			s = s[batchSize:]
		}
	}

	return batches
}
