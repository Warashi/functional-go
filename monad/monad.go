package monad

type Monad[T, U, MT, MU any] interface {
	Unitter[U, MU]
	Binder[T, MT, MU]
}

type AdditiveMonad[T, U, MT, MU any] interface {
	Monad[T, U, MT, MU]
	Zeroer[MU]
	Pluser[MT]
}

type Unitter[T, MT any] interface {
	Unit(value T) MT
}

type Binder[T, MT, MU any] interface {
	Bind(m MT, f func(T) MU) MU
}

type Zeroer[T any] interface {
	Zero() T
}

type Pluser[MT any] interface {
	Plus(MT, MT) MT
}

func Map[MU, T, U, MT any, Impl Monad[T, U, MT, MU]](impl Impl, src MT, f func(T) U) MU {
	return impl.Bind(src, func(value T) MU { return impl.Unit(f(value)) })
}

func FlatMap[T, MU, MT any, Impl Binder[T, MT, MU]](impl Impl, src MT, f func(T) MU) MU {
	return impl.Bind(src, f)
}

func Filter[T, MT any, Impl AdditiveMonad[T, T, MT, MT]](impl Impl, src MT, f func(T) bool) MT {
	return impl.Bind(src, func(value T) MT {
		if f(value) {
			return impl.Unit(value)
		}
		return impl.Zero()
	})
}

func Do[U, MU, T, MT any, Impl AdditiveMonad[T, U, MT, MU]](impl Impl, src MT, f func(T)) {
	impl.Bind(src, func(value T) MU {
		f(value)
		return impl.Zero()
	})
}
