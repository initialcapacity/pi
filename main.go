package main

import (
	"fmt"
	"iter"
	"math/big"
	"math/rand/v2"
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
func EstimatePi(points iter.Seq[Point], keepGoing func(uint64) bool) (pi float64, iterations uint64) {
	var totalPoints, insidePoints uint64

	for point := range points {
		totalPoints++
		if point.InsideUnitCircle() {
			insidePoints++
		}
		if !keepGoing(totalPoints) {
			break
		}
	}

	return DivideUint64(insidePoints, totalPoints) * 4, totalPoints
}

func main() {
	start := time.Now()
	points := GeneratePoints()
	pi, iterations := EstimatePi(points, func(iteration uint64) bool {
		if iteration%10_000_000 == 0 {
			print(".")
		}
		return iteration < 1_000_000_000
	})

	elapsed := time.Since(start)
	fmt.Printf("\nπ ≈ %.12f (n=%d, %s)\n", pi, iterations, elapsed)
}
