package cli_test

import (
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
	"github.com/initialcapacity/pi/internal/cli"
)

func TestParseCommandLineArgs(t *testing.T) {
	duration, numberOfWorkers := cli.ParseCommandLineArgs([]string{"-d", "30", "-n", "15"})
	assert.Equal(t, 30, duration)
	assert.Equal(t, 15, numberOfWorkers)

	defaultDuration, defaultNumberOfWorkers := cli.ParseCommandLineArgs([]string{})
	assert.Equal(t, 20, defaultDuration)
	assert.GreaterThanOrEqualTo(t, defaultNumberOfWorkers, 1)
}

func TestPrettyPrint(t *testing.T) {
	assert.Equal(t, "2,300", cli.PrettyPrint(2300))
	assert.Equal(t, "22,300", cli.PrettyPrint(22300))
	assert.Equal(t, "222,300", cli.PrettyPrint(222300))
}

func TestReport(t *testing.T) {
	assert.Equal(t, "π ≈ 3.141592600000 (10,000,000 iterations)", cli.Report(3.1415926, 10_000_000))
	assert.Equal(t, "π ≈ 3.000000000000 (4 iterations)", cli.Report(3, 4))
}
