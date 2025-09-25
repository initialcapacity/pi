package pi

import (
	"context"
	"iter"
	"sync"
	"sync/atomic"

	"github.com/tygern/pi/pkg/precise"
	"github.com/tygern/pi/pkg/unitsquare"
)


func insideUnitCircle(p unitsquare.Point) bool {
	return p.X*p.X+p.Y*p.Y <= 1
}

func countInsideUnitCircle(points iter.Seq[unitsquare.Point]) (total, inside uint64) {
	for point := range points {
		total++
		if insideUnitCircle(point) {
			inside++
		}
	}
	return total, inside
}

func Estimate(ctx context.Context, numberOfWorkers int) (pi float64, iterations uint64) {
	var totalPoints, insidePoints atomic.Uint64

	wg := sync.WaitGroup{}
	for _ = range numberOfWorkers {
		wg.Go(func() {
			points := unitsquare.GeneratePoints(ctx)
			total, inside := countInsideUnitCircle(points)
			totalPoints.Add(total)
			insidePoints.Add(inside)
		})
	}
	wg.Wait()

	return precise.DivideUint64(insidePoints.Load(), totalPoints.Load()) * 4, totalPoints.Load()
}
