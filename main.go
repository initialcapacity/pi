package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand/v2"
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
// It returns a channel of Point structs and a cancel function which stops the generation.
func GeneratePoints() (<-chan Point, func()) {
	points := make(chan Point)
	done := make(chan struct{})

	go func() {
		defer close(points)
		for {
			select {
			case <-done:
				return
			case points <- Point{X: rand.Float64(), Y: rand.Float64()}:
			}
		}
	}()

	return points, func() { close(done) }
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
// It stops when it reaches math.MaxUint64 iterations or when the source of points is exhausted.
func EstimatePi(points <-chan Point, reportProgress func(uint64)) (pi float64, iterations uint64) {
	var totalPoints, insidePoints uint64

	for point := range points {
		totalPoints++
		if point.InsideUnitCircle() {
			insidePoints++
		}
		reportProgress(totalPoints)
		if totalPoints == math.MaxUint64 {
			break
		}
	}

	return DivideUint64(insidePoints, totalPoints) * 4, totalPoints
}

func main() {
	points, cancelGeneration := GeneratePoints()

	pi, iterations := EstimatePi(points, func(iteration uint64) {
		if iteration%1_000_000 == 0 {
			print(".")
		}
		if iteration == 100_000_000 {
			cancelGeneration()
		}
	})

	fmt.Printf("\nπ ≈ %.12f (n=%d)\n", pi, iterations)
}
