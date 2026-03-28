package easy

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, val := range m {
		values = append(values, val)
	}

	return values
}

func SliceToMap[T any, K comparable, V any](s []T, toMapFunc func(e T) (K, V)) map[K]V {
	m := make(map[K]V, len(s))

	for _, element := range s {
		key, val := toMapFunc(element)
		m[key] = val
	}

	return m
}
