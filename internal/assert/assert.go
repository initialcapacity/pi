package assert

import (
	"cmp"
	"context"
	"strings"
	"testing"
	"time"
)

func Equal[V comparable](t testing.TB, expected, actual V) {
	t.Helper()

	if expected != actual {
		t.Fatalf(`Test assertion failed.
expected: %v
actual: %v`, expected, actual)
	}
}

func GreaterThanOrEqualTo[V cmp.Ordered](t testing.TB, first, second V) {
	t.Helper()

	if first < second {
		t.Fatalf(`Test assertion failed.
expected %v to be greater than or equal to %v`, first, second)
	}
}

func NoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf(`Test assertion failed.
expected error '%v' to be nil`, err)
	}
}

func ContainsSubstring(t testing.TB, subject, substring string) {
	t.Helper()

	if !strings.Contains(subject, substring) {
		t.Fatalf(`Test assertion failed.
expected '%v' to contain '%v'`, subject, substring)
	}
}

type PanicConfig struct {
	Timeout time.Duration
}

type PanicOption func(*PanicConfig)

func WithTimeout(timeout time.Duration) PanicOption {
	return func(c *PanicConfig) {
		c.Timeout = timeout
	}
}

func NewPanicConfig(options ...PanicOption) *PanicConfig {
	cfg := &PanicConfig{Timeout: 100 * time.Millisecond}
	for _, opt := range options {
		opt(cfg)
	}
	return cfg
}

func Panics(t testing.TB, fn func(), options ...PanicOption) any {
	t.Helper()
	recovered := make(chan any)
	config := NewPanicConfig(options...)
	ctx, cancel := context.WithTimeout(t.Context(), config.Timeout)
	defer cancel()

	// Run in a goroutine to catch any rumtime.Goexit calls
	go func() {
		defer func() {
			recovered <- recover()
		}()
		fn()
		t.Fatal("Expected function to panic")
	}()

	select {
	case r := <-recovered:
		return r
	case <-ctx.Done():
		t.Fatalf("Timeout (limit %v) waiting for panic", config.Timeout)
		return nil
	}
}
