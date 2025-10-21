package assert_test

import (
	"errors"
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
)

func assertPasses(t *testing.T, assertion func(testing.TB)) {
	t.Helper()
	subjectTest := new(testing.T)

	assert.AllowPanic(t, func() {
		assertion(subjectTest)
	})
	if subjectTest.Failed() {
		t.Error("Expected assertion to pass, but it failed")
	}
}

func assertFails(t *testing.T, assertion func(testing.TB)) {
	t.Helper()
	subjectTest := new(testing.T)

	assert.AllowPanic(t, func() {
		assertion(subjectTest)
	})

	if !subjectTest.Failed() {
		t.Error("Expected assertion to fail, but it passed")
	}
}

func TestEquals(t *testing.T) {
	assertPasses(t, func(tb testing.TB) { assert.Equal(tb, 3, 3) })
	assertFails(t, func(tb testing.TB) { assert.Equal(tb, 3, 4) })
}

func TestContainsSubstring(t *testing.T) {
	assertPasses(t, func(tb testing.TB) {
		assert.ContainsSubstring(tb, "hello world", "world")
		assert.ContainsSubstring(tb, "hello world", "ello w")
	})
	assertFails(t, func(tb testing.TB) { assert.ContainsSubstring(tb, "hello world", "goodbye") })
}

func TestNoError(t *testing.T) {
	assertPasses(t, func(tb testing.TB) { assert.NoError(tb, nil) })
	assertFails(t, func(tb testing.TB) { assert.NoError(tb, errors.New("some error")) })
}

func TestGreaterThanOrEqualTo(t *testing.T) {
	assertPasses(t, func(tb testing.TB) {
		assert.GreaterThanOrEqualTo(tb, 5, 3)
		assert.GreaterThanOrEqualTo(tb, 3, 3)
	})
	assertFails(t, func(tb testing.TB) { assert.GreaterThanOrEqualTo(tb, 2, 5) })
}
