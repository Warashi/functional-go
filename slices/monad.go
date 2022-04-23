package slices

import "github.com/Warashi/functional-go/monad"

var (
	_ monad.AdditiveMonad[int, string, []int, []string] = MonadImpl[int, string]{}
)

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) []U {
	return []U{value}
}

func (MonadImpl[T, U]) Bind(src []T, f func(T) []U) []U {
	var result []U
	for _, v := range src {
		result = append(result, f(v)...)
	}
	return result
}

func (MonadImpl[T, U]) Zero() []U {
	return nil
}

func (MonadImpl[T, U]) Plus(a, b []T) []T {
	r := make([]T, len(a)+len(b))
	copy(r[:len(a)], a)
	copy(r[len(a):], b)
	return r
}
