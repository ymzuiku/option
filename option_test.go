package option_test

import (
	"testing"

	"github.com/ymzuiku/option"
)

type User struct {
	Name string
}

type Input struct {
	User  option.Option[*User]
	Other string
}

func TestOption(t *testing.T) {
	// Initialize environment
	t.Parallel()

	var value string

	someFuncSafe := func(input *Input) {
		input.User.IfSome(func(u *User) {
			value = u.Name
		})
	}

	// Run subtests
	t.Run("Subtest Name", func(t *testing.T) {
		someFuncSafe(&Input{
			User: option.Some(&User{
				Name: "the name",
			}),
			Other: "hello",
		})
		if value != "the name" {
			t.Fatal("need eq")
		}
	})

	t.Run("load ok", func(t *testing.T) {
		oldUser := &User{
			Name: "the name2",
		}
		user := option.Wrap(oldUser)

		v, ok := user.Value()

		if !ok {
			t.Fatal("need ok")
		}
		if v.Name != "the name2" {
			t.Fatal("need value eq")
		}
	})

	t.Run("load no ok", func(t *testing.T) {
		var oldUser *User
		user := option.Wrap(oldUser)

		v, ok := user.Value()
		if ok {
			t.Fatal("need !ok")
		}
		if v != nil {
			t.Fatal("need value nil")
		}
	})
}
