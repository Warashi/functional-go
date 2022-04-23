package slices

import (
	"github.com/Warashi/functional-go/monad"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

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
func Sorted[T constraints.Ordered](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	slices.Sort(r)
	return r
}
