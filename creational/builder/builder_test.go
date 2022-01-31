package builder_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/creational/builder"
	"github.com/stretchr/testify/assert"
)

func Test_Builder(t *testing.T) {
	// build a house with floor, roof, windows and walls
	// should be finalized with "Build()" call
	house := builder.NewSimpleHouseBuilder().WithFloor().WithRoof().WithWindows().WithWalls().Build()

	assert.Equal(t, &builder.HouseComponent{Type: builder.Floor, Name: "Simple floor"}, house.Has(builder.Floor))

	assert.Equal(t, &builder.HouseComponent{Type: builder.Walls, Name: "Simple walls"}, house.Has(builder.Walls))

	assert.Equal(t, &builder.HouseComponent{Type: builder.Windows, Name: "Simple windows"}, house.Has(builder.Windows))

	assert.Equal(t, &builder.HouseComponent{Type: builder.Roof, Name: "Simple roof"}, house.Has(builder.Roof))
}
