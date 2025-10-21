package assert

import (
	"context"
	"testing"
	"time"
)

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

func AllowPanic(t testing.TB, fn func(), options ...PanicOption) any {
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
	}()

	select {
	case r := <-recovered:
		return r
	case <-ctx.Done():
		t.Fatalf("Operation timed out (limit %v) waiting for AllowPanic", config.Timeout)
		return nil
	}
}
