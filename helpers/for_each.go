package hxhelpers

func ForEachValue[T, R any](values []T, process func(T) R) []R {
	result := make([]R, len(values), len(values))

	for ix, value := range values {
		result[ix] = process(value)
	}

	return result
}

type ParamsForEachValueWAddition[T, R any] struct {
	Values []T

	Addition func() R
	Process  func(T) R
}

func ForEachValueWAddition[T, R any](params *ParamsForEachValueWAddition[T, R]) []R {
	result := make([]R, 0, len(params.Values)+1)

	result = append(result, params.Addition())

	for _, value := range params.Values {
		result = append(result, params.Process(value))
	}

	return result
}
