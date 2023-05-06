package bench

import (
	"reflect"
)

// don't use very very slowy
func authNewIfNil[T any](value T) T {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return reflect.New(v.Type().Elem()).Interface().(T)
	}
	return value
}

func authNewIfNil2[T any](value T) T {
	if IsNil(value) {
		v := reflect.ValueOf(value)
		return reflect.New(v.Type().Elem()).Interface().(T)
	}
	return value
}

func authNewIfNil3[T any](value T, safeValue T) T {
	if IsNil(value) {
		return safeValue
	}
	return value
}
