package bench

import "reflect"

type isNiler interface {
	isNil() bool
}

// https://github.com/go-x-pkg/isnil/blob/master/isnil.go

//go:inline
func IsNil(v any) bool {
	if v == nil {
		return true
	}

	if checker, ok := v.(isNiler); ok {
		return checker.isNil()
	}

	return (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
