package easy

func ConvertNullable[Tsrc any, Tdest any](ptr *Tsrc, convertFunc func(val Tsrc) Tdest) *Tdest {
	if ptr == nil {
		return nil
	}

	convertedVal := convertFunc(*ptr)
	return &convertedVal
}

func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
