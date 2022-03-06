package set

import "github.com/Warashi/functional-go/monad"

var (
	_ monad.AdditiveMonad[int, string, Set[int], Set[string]] = MonadImpl[int, string]{}
)

type MonadImpl[T, U comparable] struct{}

func (MonadImpl[T, U]) Unit(value U) Set[U] {
	return Set[U]{value: struct{}{}}
}

func (MonadImpl[T, U]) Bind(src Set[T], f func(T) Set[U]) Set[U] {
	result := Set[U]{}
	for v := range src {
		for k := range f(v) {
			result[k] = struct{}{}
		}
	}
	return result
}

func (MonadImpl[T, U]) Zero() Set[U] {
	return Set[U]{}
}

func (MonadImpl[T, U]) Plus(a, b Set[T]) Set[T] {
	result := Set[T]{}
	for k := range a {
		result[k] = struct{}{}
	}
	for k := range b {
		result[k] = struct{}{}
	}
	return result
}
