package option

import "reflect"

type isNiler interface {
	isNil() bool
}

// https://github.com/go-x-pkg/isnil/blob/master/isnil.go
// IsNil checks if any is nil.
func isNil(v any) bool {
	if v == nil {
		return true
	}

	if checker, ok := v.(isNiler); ok {
		return checker.isNil()
	}

	return (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
