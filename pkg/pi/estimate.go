package pi

import (
	"context"
	"iter"
	"sync"
	"sync/atomic"

	"github.com/tygern/pi/pkg/precise"
	"github.com/tygern/pi/pkg/unitsquare"
)


func InsideUnitCircle(p unitsquare.Point) bool {
	return p.X*p.X+p.Y*p.Y <= 1
}

func CountInsideUnitCircle(points iter.Seq[unitsquare.Point]) (total, inside uint64) {
	for point := range points {
		total++
		if InsideUnitCircle(point) {
			inside++
		}
	}
	return total, inside
}

type Generate func(ctx context.Context) iter.Seq[unitsquare.Point]

func Estimate(ctx context.Context, generate Generate, numberOfWorkers int) (pi float64, iterations uint64) {
	var totalPoints, insidePoints atomic.Uint64

	wg := sync.WaitGroup{}
	for _ = range numberOfWorkers {
		wg.Go(func() {
			points := generate(ctx)
			total, inside := CountInsideUnitCircle(points)
			totalPoints.Add(total)
			insidePoints.Add(inside)
		})
	}
	wg.Wait()

	return precise.DivideUint64(insidePoints.Load(), totalPoints.Load()) * 4, totalPoints.Load()
}
