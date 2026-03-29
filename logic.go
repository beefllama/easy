package easy

// ConvertNullable converts a nullable value into nullable value of another type.
func ConvertNullable[Tsrc any, Tdest any](ptr *Tsrc, convertFunc func(val Tsrc) *Tdest) *Tdest {
	if ptr == nil {
		return nil
	}

	return convertFunc(*ptr)
}

// Ternary be cautious about dereferencing pointers.
func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
