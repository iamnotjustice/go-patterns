package strategy_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/behavioral/strategy"
	"github.com/stretchr/testify/assert"
)

func Test_Strategy_Struct(t *testing.T) {
	trackerHaver := strategy.NewTracker(&strategy.Haversine{})

	tr := trackerHaver.GetTimeRemaining(&strategy.Point{X: 51.5720575, Y: -0.2322095}, &strategy.Point{X: 51.5433816, Y: -0.199606}, 40)

	assert.Equal(t, int64(4), tr)

	trackerGoogle := strategy.NewTracker(&strategy.GoogleDistance{})

	tr = trackerGoogle.GetTimeRemaining(
		&strategy.Point{X: 51.5720575, Y: -0.2322095},
		&strategy.Point{X: 51.5433816, Y: -0.199606},
		40,
	)

	assert.Equal(t, int64(6), tr)
}

func Test_Strategy_Func(t *testing.T) {
	trackerHaver := strategy.NewTrackerWithFunc(strategy.HaversineCalc)

	tr := trackerHaver.GetTimeRemaining(&strategy.Point{X: 51.5720575, Y: -0.2322095}, &strategy.Point{X: 51.5433816, Y: -0.199606}, 40)

	assert.Equal(t, int64(4), tr)
}
