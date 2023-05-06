package option

type Option[T any] struct {
	v  T
	ok bool
}

//go:inline
func Some[T any](v T) Option[T] {
	return Option[T]{v: v, ok: true}
}

//go:inline
func None[T any]() Option[T] {
	var zero T
	return Option[T]{v: zero, ok: false}
}

//go:inline
func Wrap[T any](v T) Option[T] {
	return Option[T]{v: v, ok: !IsNil(v)}
}

func WrapCheck[T any](v T, ok bool) Option[T] {
	return Option[T]{v: v, ok: ok}
}

//go:inline
func (o *Option[T]) Value() (T, bool) {
	return o.v, o.ok
}

//go:inline
func (o *Option[T]) ValueOrError(err error) (T, error) {
	if !o.ok {
		var zero T
		return zero, err
	}
	return o.v, nil
}

//go:inline
func (o *Option[T]) Match(someFn func(T), noneFn func()) {
	if o.ok {
		someFn(o.v)
	} else {
		noneFn()
	}
}

//go:inline
func (o *Option[T]) IfSome(someFn func(T)) {
	if o.ok {
		someFn(o.v)
	}
}

//go:inline
func (o *Option[T]) IfNone(noneFn func()) {
	if !o.ok {
		noneFn()
	}
}

//go:inline
func (o *Option[T]) Unwrap_unsafe() T {
	if !o.ok {
		panic("Value is nil")
	}
	return o.v
}

//go:inline
func (o *Option[T]) Expect_unsafe(err error) T {
	if !o.ok {
		panic(err)
	}
	return o.v
}
