package strategy

import (
	"math"
	"net/url"
)

type Point struct {
	X float64
	Y float64
}

type Distancer interface {
	Calculate(*Point, *Point) int64
}

type Haversine struct{}

func (h *Haversine) Calculate(x, y *Point) int64 {
	const (
		earthRadiusMi = 3958 // radius of the earth in miles.
		earthRaidusKm = 6371 // radius of the earth in kilometers.
	)

	degToRad := func(d float64) float64 {
		return d * math.Pi / 180
	}

	lat1 := degToRad(x.X)
	lon1 := degToRad(x.Y)
	lat2 := degToRad(y.X)
	lon2 := degToRad(y.Y)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km := int64(math.Floor(c * earthRaidusKm))

	return km
}

type GoogleDistance struct {
	url url.URL
}

func (gd *GoogleDistance) Calculate(x, y *Point) int64 {
	// imagine a http call to Google API to get more accurate distance

	// but here we just hand over this... mock
	return 4
}

type Tracker struct {
	distanceCalc Distancer
}

func NewTracker(calc Distancer) *Tracker {
	return &Tracker{
		distanceCalc: calc,
	}
}

func (t *Tracker) GetTimeRemaining(current, finish *Point, speed int64) int64 {
	distance := t.distanceCalc.Calculate(current, finish)

	time := float64(distance) / float64(speed) * 60

	return int64(math.Floor(time))
}
