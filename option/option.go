package option

import "github.com/Warashi/functional-go/monad"

type Option[T any] struct {
	value *T
}

func (o Option[T]) IsSome() bool { return o.value != nil }
func (o Option[T]) IsNone() bool { return o.value == nil }
func (o Option[T]) OrElse(v T) T {
	if o.value == nil {
		return v
	}
	return *o.value
}

func Map[T, U comparable](from Option[T], f func(T) U) Option[U] {
	return monad.Map[Option[U]](MonadImpl[T, U]{}, from, f)
}
func FlatMap[T, U comparable](from Option[T], f func(T) Option[U]) Option[U] {
	return monad.FlatMap(MonadImpl[T, U]{}, from, f)
}
func Filter[T comparable](from Option[T], f func(T) bool) Option[T] {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T comparable](from Option[T], f func(T)) {
	monad.Do[struct{}, Option[struct{}]](MonadImpl[T, struct{}]{}, from, f)
}
