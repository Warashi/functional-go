package sets

import (
	"github.com/Warashi/functional-go/monad"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](values ...T) Set[T] {
	s := make(Set[T], len(values))
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func Map[T, U comparable](from Set[T], f func(T) U) Set[U] {
	return monad.Map[Set[U]](MonadImpl[T, U]{}, from, f)
}
func FlatMap[T, U comparable](from Set[T], f func(T) Set[U]) Set[U] {
	return monad.FlatMap(MonadImpl[T, U]{}, from, f)
}
func Filter[T comparable](from Set[T], f func(T) bool) Set[T] {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T comparable](from Set[T], f func(T)) {
	monad.Do[struct{}, Set[struct{}]](MonadImpl[T, struct{}]{}, from, f)
}
func AsSlice[T comparable](from Set[T]) []T {
	s := make([]T, 0, len(from))
	ForEach(from, func(t T) {
		s = append(s, t)
	})
	return s
}
func AsSliceSorted[T constraints.Ordered](from Set[T]) []T {
	s := AsSlice(from)
	slices.Sort(s)
	return s
}
func AsSliceSortedFunc[T comparable](from Set[T], less func(a, b T) bool) []T {
	s := AsSlice(from)
	slices.SortFunc(s, less)
	return s
}
