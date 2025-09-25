package main

import (
	"context"
	"flag"
	"fmt"
	"iter"
	"math/big"
	"math/rand/v2"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Point struct {
	X, Y float64
}

func (p Point) InsideUnitCircle() bool {
	return p.X*p.X+p.Y*p.Y <= 1
}

func GeneratePoints(ctx context.Context) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for {
			ok := yield(Point{X: rand.Float64(), Y: rand.Float64()})
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}
}

func DivideUint64(numerator, denominator uint64) float64 {
	bigNumerator := new(big.Int).SetUint64(numerator)
	bigDenominator := new(big.Int).SetUint64(denominator)
	ratio := new(big.Rat).SetFrac(bigNumerator, bigDenominator)

	approx, _ := ratio.Float64()
	return approx
}

func CountInsideUnitCircle(ctx context.Context, points iter.Seq[Point]) (total, inside uint64) {
	for point := range points {
		total++
		if point.InsideUnitCircle() {
			inside++
		}
	}
	return total, inside
}

func EstimatePi(ctx context.Context, numberOfWorkers int) (pi float64, iterations uint64) {
	var totalPoints, insidePoints atomic.Uint64

	wg := sync.WaitGroup{}
	for _ = range numberOfWorkers {
		wg.Go(func() {
			total, inside := CountInsideUnitCircle(ctx, GeneratePoints(ctx))
			totalPoints.Add(total)
			insidePoints.Add(inside)
		})
	}
	wg.Wait()

	return DivideUint64(insidePoints.Load(), totalPoints.Load()) * 4, totalPoints.Load()
}

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

func main() {
	duration, numberOfWorkers := ParseCommandLineArgs()

	ctx, cancel := SigtermTimeoutContext(time.Duration(duration) * time.Second)
	defer cancel()
	done := ExecutionTimer("Estimation")
	defer done()

	p := message.NewPrinter(language.English)
	pi, iterations := EstimatePi(ctx, numberOfWorkers)
	p.Printf("π ≈ %.12f (%d iterations)\n", pi, iterations)
}
