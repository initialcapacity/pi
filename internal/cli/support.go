package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func ParseCommandLineArgs() (duration, numberOfWorkers int) {
	flag.IntVar(&duration, "d", 20, "execution duration in seconds")
	flag.IntVar(&numberOfWorkers, "n", runtime.NumCPU(), "number of workers")
	flag.Parse()
	fmt.Printf("Running for %d seconds with %d workers\n", duration, numberOfWorkers)
	return duration, numberOfWorkers
}

func Report(pi float64, iterations int) {
	printer := message.NewPrinter(language.English)
	printer.Printf("π ≈ %.12f (%d iterations)\n", pi, iterations)
}
