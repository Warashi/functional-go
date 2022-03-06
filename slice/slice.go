package slice

import "github.com/Warashi/functional-go/monad"

func Map[F, T any](from []F, f func(F) T) []T {
	return monad.Map[[]T](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from []F, f func(F) []T) []T {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T any](from []T, f func(T) bool) []T {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T any](from []T, f func(T)) {
	monad.Do[struct{}, []struct{}](MonadImpl[T, struct{}]{}, from, f)
}
