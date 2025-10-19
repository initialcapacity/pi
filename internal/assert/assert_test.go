package assert_test

import (
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
)

func assertPasses(t *testing.T, assertion func(testing.TB)) {
	subjectTest := new(testing.T)

	assertion(subjectTest)

	if subjectTest.Failed() {
		t.Error("Expected assertion to pass, but it failed")
	}
}

func assertFails(t *testing.T, assertion func(testing.TB)) {
	subjectTest := new(testing.T)

	assert.Panics(t, func() {
		assertion(subjectTest)
	})

	if !subjectTest.Failed() {
		t.Error("Expected assertion to fail, but it passed")
	}
}

func TestEquals(t *testing.T) {
	assertPasses(t, func(tb testing.TB) {
		assert.Equal(tb, 3, 3)
	})
	assertFails(t, func(tb testing.TB) {
		assert.Equal(tb, 3, 4)
	})
}
