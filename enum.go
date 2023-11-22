package enum

type enum[T any] struct {
	slice []T
}

func Of[T any](items []T) *enum[T] {
	return &enum[T]{
		slice: items,
	}
}

func (e enum[T]) ToSlice() []T {
	return e.slice
}

// Slice drops elements before first (zero-base), then it tates elements until element last (inclusively).
func (e enum[T]) Slice(first, last int) enum[T] {
	return enum[T]{
		slice: e.slice[first : last+1],
	}
}

func (e enum[T]) Filter(f func(elem T) bool) enum[T] {
	result := make([]T, 0)
	for _, v := range e.slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return enum[T]{
		slice: result,
	}
}

func (e enum[T]) Find(f func(elem T) bool) (item *T, ok bool) {
	for _, v := range e.slice {
		if f(v) {
			return &v, true
		}
	}
	return nil, false
}

func (e enum[T]) Map(f func(item T) T) enum[T] {
	result := make([]T, 0)
	for _, v := range e.slice {
		result = append(result, f(v))
	}
	return enum[T]{
		slice: result,
	}
}

func (e enum[T]) Each(f func(item T)) enum[T] {
	for _, v := range e.slice {
		f(v)
	}
	return e
}

func (e enum[T]) Reduce(a any, f func(item T, a any) any) any {
	result := a
	for _, v := range e.slice {
		result = f(v, result)
	}
	return result
}

func (e enum[T]) ReduceWithIndex(a any, f func(i int, item T, a any) any) any {
	result := a
	for i, v := range e.slice {
		result = f(i, v, result)
	}
	return result
}
