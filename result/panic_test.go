package result_test

import (
	"errors"
	"fmt"
	"testing"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func DividePanic(x, y float64) float64 {
	_ = fibonacci(20)
	if y == 0 {
		panic(errors.New("division by zero"))
	}
	return x / y
}

func D1() float64 {
	return D2()
}

func D2() float64 {
	return D3()
}

func D3() float64 {
	return DividePanic(10, 2)
}

func E1() (float64, error) {
	return E2()
}

func E2() (float64, error) {
	return E3()
}

func E3() (float64, error) {
	return DivideSafely(10, 2)
}

func DivideSafely(x, y float64) (float64, error) {
	_ = fibonacci(20)
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

func BenchmarkDividePanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		D1()
	}
}

func BenchmarkDivideSafely(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := E1(); err != nil {
			fmt.Println(err)
		}
	}
}
