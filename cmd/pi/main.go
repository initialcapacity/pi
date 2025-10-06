package main

import (
	"fmt"
	"os"
	"time"

	"github.com/initialcapacity/pi/internal/cli"
	"github.com/initialcapacity/pi/pkg/pi"
	"github.com/initialcapacity/pi/pkg/unitsquare"
)

func main() {
	duration, numberOfWorkers := cli.ParseCommandLineArgs(os.Args[1:])

	ctx, cancel := cli.SigtermTimeoutContext(time.Duration(duration) * time.Second)
	defer cancel()
	done := cli.ExecutionTimer("Estimation")
	defer done()

	pi, iterations := pi.Estimate(ctx, unitsquare.GeneratePoints, numberOfWorkers)
	fmt.Println(cli.Report(pi, int(iterations)))
}
