package unitsquare

import (
	"context"
	"iter"
	"math/rand/v2"
)

type Point struct {
	X, Y float64
}

func GeneratePoints(ctx context.Context) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for {
			for _ = range 1000 {
				ok := yield(Point{X: rand.Float64(), Y: rand.Float64()})
				if !ok {
					return
				}
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}
}
