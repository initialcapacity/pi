package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand/v2"
	"time"
)

type Point struct {
	X, Y float64
}

func (p Point) InsideUnitCircle() bool {
	return p.X*p.X+p.Y*p.Y <= 1
}

func GeneratePoints(ctx context.Context) <-chan Point {
	points := make(chan Point, 1_000)
	go func() {
		defer close(points)
		for {
			select {
			case <-ctx.Done():
				return
			case points <- Point{X: rand.Float64(), Y: rand.Float64()}:
			}
		}
	}()
	return points
}

func BigRatio(numerator, denominator uint64) float64 {
	bigNumerator := new(big.Int).SetUint64(numerator)
	bigDenominator := new(big.Int).SetUint64(denominator)
	ratio := new(big.Rat).SetFrac(bigNumerator, bigDenominator)

	approx, _ := ratio.Float64()
	return approx
}

func EstimatePi(points <-chan Point, progressTracker func(uint64)) (pi float64, iterations uint64) {
	var totalPoints, insidePoints uint64

	for point := range points {
		totalPoints++
		if point.InsideUnitCircle() {
			insidePoints++
		}
		progressTracker(totalPoints)
		if totalPoints == math.MaxUint64 {
			break
		}
	}

	return BigRatio(insidePoints, totalPoints) * 4, totalPoints
}

func PrintProgress(currentIteration uint64) {
	if currentIteration%1_000_000 == 0 {
		print(".")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	points := GeneratePoints(ctx)
	pi, iterations := EstimatePi(points, PrintProgress)

	fmt.Printf("\nπ ≈ %.12f (n=%d)\n", pi, iterations)
}
