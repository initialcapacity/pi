package unitsquare_test

import (
	"context"
	"slices"
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
	"github.com/initialcapacity/pi/pkg/unitsquare"
)

func TestGeneratePoints(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	points := unitsquare.GeneratePoints(ctx)
	cancel()
	result := slices.Collect(points)

	assert.GreaterThanOrEqualTo(t, len(result), 1_000)
}
