package option_test

import (
	"testing"

	"github.com/ymzuiku/option"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		check := func(user *User) bool {
			_ = fibonacci(5)
			return user == nil
		}
		_ = check(&User{})
		var user User
		_ = check(&user)
	}
}

func BenchmarkDivIsNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		check := func(user *User) bool {
			_ = fibonacci(5)
			return option.IsNil(user)
		}
		_ = check(&User{})
		var user User
		_ = check(&user)
	}
}
