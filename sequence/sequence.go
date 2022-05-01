package sequence

import (
	"github.com/Warashi/functional-go/monad"
)

var (
	_ Sequence[any] = (*SliceSequence[any])(nil)
)

type SliceSequence[T any] struct {
	cursor int
	base   []T
}

func (s *SliceSequence[T]) Next() bool {
	if s.cursor+1 >= len(s.base) {
		return false
	}
	s.cursor++
	return true
}

func (s *SliceSequence[T]) Value() T {
	if s.cursor < 0 {
		var zero T
		return zero
	}
	return s.base[s.cursor]
}

func Of[T any](value ...T) Sequence[T] {
	return &SliceSequence[T]{
		cursor: -1,
		base:   value,
	}
}

type Sequence[T any] interface {
	Next() bool
	Value() T
}

func Map[F, T any](from Sequence[F], f func(F) T) Sequence[T] {
	return monad.Map[Sequence[T]](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from Sequence[F], f func(F) Sequence[T]) Sequence[T] {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T any](from Sequence[T], f func(T) bool) Sequence[T] {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func Flatten[T any](s Sequence[Sequence[T]]) Sequence[T] {
	return FlatMap(s, func(v Sequence[T]) Sequence[T] { return v })
}

// ForEach
// 遅延評価しているため、終端処理はmonadで実装できない
func ForEach[T any](s Sequence[T], f func(T)) {
	for s.Next() {
		f(s.Value())
	}
}
func Collect[T any](s Sequence[T]) []T {
	var v []T
	ForEach(s, func(value T) { v = append(v, value) })
	return v
}
