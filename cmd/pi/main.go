package main

import (
	"fmt"
	"time"

	"github.com/tygern/pi/internal/cli"
	"github.com/tygern/pi/pkg/pi"
	"github.com/tygern/pi/pkg/unitsquare"
)

func main() {
	duration, numberOfWorkers := cli.ParseCommandLineArgs()

	ctx, cancel := cli.SigtermTimeoutContext(time.Duration(duration) * time.Second)
	defer cancel()
	done := cli.ExecutionTimer("Estimation")
	defer done()

	pi, iterations := pi.Estimate(ctx, unitsquare.GeneratePoints, numberOfWorkers)
	fmt.Println(cli.Report(pi, int(iterations)))
}
