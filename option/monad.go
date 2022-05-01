package option

import "github.com/Warashi/functional-go/monad"

var (
	_ monad.AdditiveMonad[int, string, Option[int], Option[string]] = MonadImpl[int, string]{}
)

type MonadImpl[T, U comparable] struct{}

func (MonadImpl[T, U]) Unit(value U) Option[U] {
	return Option[U]{value: &value}
}

func (MonadImpl[T, U]) Bind(src Option[T], f func(T) Option[U]) Option[U] {
	if src.IsSome() {
		return f(*src.value)
	}
	return Option[U]{}
}

func (MonadImpl[T, U]) Zero() Option[U] {
	return Option[U]{}
}

func (MonadImpl[T, U]) Plus(a, b Option[T]) Option[T] {
	if a.IsSome() {
		return a
	}
	return b
}
