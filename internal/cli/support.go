package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func ExecutionTimer(description string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %s\n", description, elapsed)
	}
}

func SigtermTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	sigtermChannel := make(chan os.Signal, 1)
	signal.Notify(sigtermChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigtermChannel
		cancel()
	}()
	return ctx, cancel
}

func ParseCommandLineArgs(arguments []string) (duration, numberOfWorkers int) {
	flagSet := flag.FlagSet{}
	flagSet.IntVar(&duration, "d", 20, "execution duration in seconds")
	flagSet.IntVar(&numberOfWorkers, "n", runtime.NumCPU(), "number of workers")
	flagSet.Parse(arguments)

	fmt.Printf("Running for %d seconds with %d workers\n", duration, numberOfWorkers)
	return duration, numberOfWorkers
}

func PrettyPrint(integer int) string {
	stringRepresentation := strconv.Itoa(integer)
	var formatted strings.Builder
	for i, r := range stringRepresentation {
		if i > 0 && (len(stringRepresentation)-i)%3 == 0 {
			formatted.WriteRune(',')
		}
		formatted.WriteRune(r)
	}

	return formatted.String()
}

func Report(pi float64, iterations int) string {
	return fmt.Sprintf("π ≈ %.12f (%s iterations)", pi, PrettyPrint(iterations))
}
