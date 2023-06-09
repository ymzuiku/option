# option

Package `option` provides an implementation of the `Option` type, which is used to indicate that a value may or may not be present. This is a common pattern in many programming languages, and is often used to avoid the need for null checks or other error handling code.

- Without using reflection

## Installation

```
go get github.com/ymzuiku/option
```

## Usage

To use `option`, import it into your Go code:

```go
import "github.com/ymzuiku/option"
```

You can then create `Option` values using the `Some`, `None`, or `Wrap` functions.

Case 1:

```go
type User struct {
  Name string
}

type Input struct {
  User option.Option[*User]
  Other string
}


func SomeFuncSafe(input *Input) {
  input.User.IfSome(func(u *User) {
    // safe
    fmt.Println(u)
  })
}

func main(){
  SomeFuncSafe(&Input{
    User: option.Some(&User{
      Name: "the name",
    }),
    Other: "hello",
  })
}

```

Case 2, no change old Input API:

```go
type User struct {
  Name string
}

type Input struct {
  User *User
  Other string
}

func SomeFuncUnsafe(input *Input) {
  // Oops!! I forgot to check for nil.
  fmt.Println(input.User.Name)
}

func SomeFuncSafe(input *Input) {
  user := option.Wrap(input.User, input.User != nil)
  user.IfSome(func(u *User) {
    // safe
    fmt.Println(u)
  })
}


```

You can retrieve the value of an `Option` using the `Value` method:

```go
value, ok := opt1.Value()
if ok {
    // Do something with value
} else {
    // Value is not present
}
```

Alternatively, you can retrieve the value or return an error using the `ValueOrError` method:

```go
value, err := opt1.ValueOrError(errors.New("value not present"))
if err != nil {
    // Handle error
} else {
    // Do something with value
}
```

You can also use the `Match`, `IfSome`, and `IfNone` methods to execute different code paths depending on whether the value is present:

```go
opt1.Match(func(value int) {
    // Do something with value
}, func() {
    // Value is not present
})

opt1.IfSome(func(value int) {
    // Do something with value
})

opt2.IfNone(func() {
    // Value is not present
})
```

Finally, you can use the `Unwrap_unsafe` and `Expect_unsafe` methods to retrieve the value of an `Option` directly, without checking whether the value is present. These methods will panic if the value is not present:

```go
// Retrieve the value of an Option without checking whether it is present
value := opt1.Unwrap_unsafe()

// Retrieve the value of an Option and panic with a custom error message if the value is not present
value := opt2.Expect_unsafe(errors.New("value not present"))
```

## License

`option` is released under the MIT License. See [LICENSE](https://github.com/ymzuiku/option/blob/main/LICENSE) for more information.
