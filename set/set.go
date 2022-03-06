package set

import "github.com/Warashi/functional-go/monad"

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func Map[F, T comparable](from Set[F], f func(F) T) Set[T] {
	return monad.Map[Set[T]](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T comparable](from Set[F], f func(F) Set[T]) Set[T] {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T comparable](from Set[T], f func(T) bool) Set[T] {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T comparable](from Set[T], f func(T)) {
	monad.Do[struct{}, Set[struct{}]](MonadImpl[T, struct{}]{}, from, f)
}
