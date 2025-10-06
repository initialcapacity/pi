package pi_test

import (
	"context"
	"iter"
	"slices"
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
	"github.com/initialcapacity/pi/pkg/pi"
	"github.com/initialcapacity/pi/pkg/unitsquare"
)

func TestInsideUnitCircle(t *testing.T) {
	assert.Equal(t, true, pi.InsideUnitCircle(unitsquare.Point{X: .5, Y: .5}))
	assert.Equal(t, false, pi.InsideUnitCircle(unitsquare.Point{X: .9, Y: .5}))
}

func TestCountInsideUnitCircle(t *testing.T) {
	points := slices.Values([]unitsquare.Point{
		{X: .5, Y: .5},
		{X: .9, Y: .5},
		{X: 1, Y: 1},
		{X: .1, Y: .1},
		{X: .2, Y: .1},
	})

	total, inside := pi.CountInsideUnitCircle(points)

	assert.Equal(t, 5, total)
	assert.Equal(t, 3, inside)
}

func TestEstimatePi(t *testing.T) {
	generate := func(ctx context.Context) iter.Seq[unitsquare.Point] {
		return slices.Values([]unitsquare.Point{
			{X: .5, Y: .5},
			{X: .9, Y: .5},
			{X: 1, Y: 1},
			{X: .1, Y: .1},
			{X: .2, Y: .1},
		})
	}

	pi, iterations := pi.Estimate(t.Context(), generate, 2)

	assert.Equal(t, 10, iterations)
	assert.Equal(t, 2.4, pi)
}
