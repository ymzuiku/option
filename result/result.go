package result

import "fmt"

type Result[T any] struct {
	v   T
	err error
}

//go:inline
func Ok[T any](v T) Result[T] {
	return Result[T]{v: v, err: nil}
}

//go:inline
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{v: zero, err: err}
}

//go:inline
func Warp[T any](v T, err error) Result[T] {
	return Result[T]{v: v, err: err}
}

//go:inline
func (o *Result[T]) Value() (T, error) {
	return o.v, o.err
}

//go:inline
func (o *Result[T]) ValueError(err error) (T, error) {
	if o.err != nil {
		return o.v, fmt.Errorf("%w, %w", o.err, err)
	}
	return o.v, o.err
}

//go:inline
func (o *Result[T]) Match(okFn func(T), errFn func(error)) {
	if o.err == nil {
		okFn(o.v)
	} else {
		errFn(o.err)
	}
}

//go:inline
func (o *Result[T]) IfOk(okFn func(T)) {
	if o.err == nil {
		okFn(o.v)
	}
}

//go:inline
func (o *Result[T]) IfErr(errFn func(error)) {
	if o.err != nil {
		errFn(o.err)
	}
}

//go:inline
func (o *Result[T]) Unwrap_unsafe() T {
	if o.err != nil {
		panic(o.err)
	}
	return o.v
}

//go:inline
func (o *Result[T]) Expect_unsafe(err error) T {
	if o.err != nil {
		panic(fmt.Errorf("%w, %w", o.err, err))
	}
	return o.v
}
