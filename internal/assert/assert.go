package assert

import (
	"cmp"
	"strings"
	"testing"
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
