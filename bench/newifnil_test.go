package bench

import (
	"testing"
)

type User struct {
	Name string
}

func TestNewIfNil(t *testing.T) {
	// Run subtests
	t.Run("new if nil", func(t *testing.T) {
		{
			var user User
			name := authNewIfNil(&user).Name
			if name != "" {
				t.Fatal("need load name")
			}
		}
		{
			user := &User{
				Name: "the name",
			}
			name := authNewIfNil(user).Name
			if name != "the name" {
				t.Fatal("need load name")
			}
		}
	})
}

func BenchmarkNewIfNewReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {
			var user User
			_ = authNewIfNil(&user).Name
		}()
	}
}

func BenchmarkNewIfNewReflect2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {
			var user User
			_ = authNewIfNil2(&user).Name
		}()
	}
}

func BenchmarkNewIfNewReflect3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {
			var user User
			_ = authNewIfNil3(&user, &User{}).Name
		}()
	}
}

func newNil(user *User) *User {
	if user == nil {
		return &User{}
	}
	return user
}

func BenchmarkNewIfNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dog User
		_ = newNil(&dog).Name
	}
}
