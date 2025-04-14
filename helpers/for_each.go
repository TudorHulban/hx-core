package hxhelpers

func ForEachValue[T, R any](values []T, process func(T) R) []R {
	result := make([]R, len(values), len(values))

	for ix, value := range values {
		result[ix] = process(value)
	}

	return result
}
