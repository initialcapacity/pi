package cli_test

import (
	"testing"

	"github.com/tygern/pi/internal/assert"
	"github.com/tygern/pi/internal/cli"
)

func TestPrettyPrint(t *testing.T) {
	assert.Equal(t, "2,300", cli.PrettyPrint(2300))
	assert.Equal(t, "22,300", cli.PrettyPrint(22300))
	assert.Equal(t, "222,300", cli.PrettyPrint(222300))
}
