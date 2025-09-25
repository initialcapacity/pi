package main

import (
	"flag"
	"fmt"
	"iter"
	"math/big"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Point represents a point in 2D space.
type Point struct {
	X, Y float64
}

// InsideUnitCircle checks if a point is inside the unit circle.
func (p Point) InsideUnitCircle() bool {
	return p.X*p.X+p.Y*p.Y <= 1
}

// GeneratePoints generates random points within the unit square.
// It returns an iterable sequence of Point structs.
func GeneratePoints() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for {
			ok := yield(Point{X: rand.Float64(), Y: rand.Float64()})
			if !ok {
				return
			}
		}
	}
}

// DivideUint64 divides two uint64 numbers and returns the result as a float64.
func DivideUint64(numerator, denominator uint64) float64 {
	bigNumerator := new(big.Int).SetUint64(numerator)
	bigDenominator := new(big.Int).SetUint64(denominator)
	ratio := new(big.Rat).SetFrac(bigNumerator, bigDenominator)

	approx, _ := ratio.Float64()
	return approx
}

// EstimatePi estimates the value of π using Monte Carlo simulation.
// It returns the estimated value of π and the number of iterations performed.
// It stops when the keepGoing parameter returns false.
func EstimatePi(points iter.Seq[Point], keepGoing func() bool) (pi float64, iterations uint64) {
	var totalPoints, insidePoints uint64

	for point := range points {
		totalPoints++
		if point.InsideUnitCircle() {
			insidePoints++
		}
		if !keepGoing() {
			break
		}
	}

	return DivideUint64(insidePoints, totalPoints) * 4, totalPoints
}

// ExecutionTime measures time between when it is called and when the function it returns is called
// and prints the result.
func ExecutionTime(description string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %s\n", description, elapsed)
	}
}

func RunFor(delay time.Duration) func() bool {
	sigtermChannel := make(chan os.Signal, 1)
	signal.Notify(sigtermChannel, os.Interrupt, syscall.SIGTERM)
	timer := time.NewTimer(delay)
	cancel := make(chan interface{})
	go func() {
		select {
		case <-sigtermChannel:
		case <-timer.C:
		}
		close(sigtermChannel)
		timer.Stop()
		close(cancel)
	}()
	
	return func() bool {
		select {
		case <-cancel:
			return false
		default:
			return true
		}
	}
}

func main() {
	var duration int
	flag.IntVar(&duration, "n", 20, "execution duration in seconds")
	flag.Parse()

	defer ExecutionTime("Estimation")()

	points := GeneratePoints()
	pi, iterations := EstimatePi(points, RunFor(time.Duration(duration) * time.Second))

	fmt.Printf("π ≈ %.12f (n=%d)\n", pi, iterations)
}
