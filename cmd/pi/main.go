package main

import (
	"time"

	"github.com/tygern/pi/internal/cli"
	"github.com/tygern/pi/pkg/pi"
)

func main() {
	duration, numberOfWorkers := cli.ParseCommandLineArgs()

	ctx, cancel := cli.SigtermTimeoutContext(time.Duration(duration) * time.Second)
	defer cancel()
	done := cli.ExecutionTimer("Estimation")
	defer done()

	pi, iterations := pi.Estimate(ctx, numberOfWorkers)
	cli.Report(pi, int(iterations))
}
