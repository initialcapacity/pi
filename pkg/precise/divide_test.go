package precise_test

import (
	"testing"

	"github.com/tygern/pi/internal/assert"
	"github.com/tygern/pi/pkg/precise"
)

func TestDivideUint64(t *testing.T) {
	assert.Equal(t, 2, precise.DivideUint64(8_000_000_000_000_000_000, 4_000_000_000_000_000_000))
}
