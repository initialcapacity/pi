package assert

import (
	"cmp"
	"testing"
)

func Equal[V comparable](t testing.TB, expected, actual V) {
	t.Helper()

	if expected != actual {
		t.Errorf(`Test assertion failed.
expected: %v
actual: %v`, expected, actual)
	}
}

func GreaterThanOrEqualTo[V cmp.Ordered](t testing.TB, first, second V) {
	t.Helper()

	if first < second {
		t.Errorf(`Test assertion failed.
expected %v to be greater than or equal to %v`, first, second)
	}
}
