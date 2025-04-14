package hxhelpers

func Coalesce[T comparable](first, second T) T {
	var zeroValue T

	if first == zeroValue {
		return second
	}

	return first
}

func CoalesceSlices[T comparable](first, second []T) []T {
	if len(first) == 0 {
		return second
	}

	return first
}
