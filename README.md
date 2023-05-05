# nilsafe

Package `nilsafe` provides an implementation of the `Option` type, which is used to indicate that a value may or may not be present. This is a common pattern in many programming languages, and is often used to avoid the need for null checks or other error handling code.

## Installation

```
go get github.com/ymzuiku/nilsafe
```

## Usage

To use `nilsafe`, import it into your Go code:

```go
import "github.com/ymzuiku/nilsafe"
```

You can then create `Option` values using the `Some`, `None`, or `Wrap` functions:

```go
// Create an Option with a value
opt1 := nilsafe.Some(42)

// Create an Option with no value
opt2 := nilsafe.None()

// Create an Option with a value and a flag indicating whether the value is present
opt3 := nilsafe.Wrap("hello", true)
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

`nilsafe` is released under the MIT License. See [LICENSE](https://github.com/ymzuiku/nilsafe/blob/main/LICENSE) for more information.
