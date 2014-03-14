package geo

import (
	"math"
)

// Represents a Physical Point in geographic notation [lat, lng].
type Point struct {
	Lat float64
	Lng float64
}

const (
	// According to Wikipedia, the Earth's radius is about 6,371km
	EARTH_RADIUS = 6371
)

// Returns a new Point populated by the passed in latitude (Lat) and longitude (lng) values.
func NewPoint(Lat float64, lng float64) *Point {
	return &Point{Lat: Lat, Lng: lng}
}

/*
// Returns Point p's Latitude.
func (p *Point) Lat() float64 {
	return p.Lat
}

// Returns Point p's longitude.
func (p *Point) Lng() float64 {
	return p.Lng
}
*/

// Returns a Point popuLated with the lat and lng coordinates of transposing the origin point the distance (in meters) supplied by the compass bearing (in degrees) supplied.
// Original Implementation from: http://www.movable-type.co.uk/scripts/Latlong.html
func (p *Point) PointAtDistanceAndBearing(dist float64, bearing float64) *Point {

	dr := dist / EARTH_RADIUS

	bearing = (bearing * (math.Pi / 180.0))

	lat1 := (p.Lat * (math.Pi / 180.0))
	lng1 := (p.Lng * (math.Pi / 180.0))

	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(bearing)

	lat2 := math.Asin(lat2_part1 + lat2_part2)

	lng2_part1 := math.Sin(bearing) * math.Sin(dr) * math.Cos(lat1)
	lng2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))

	lng2 := lng1 + math.Atan2(lng2_part1, lng2_part2)
	lng2 = math.Mod((lng2+3*math.Pi), (2*math.Pi)) - math.Pi

	lat2 = lat2 * (180.0 / math.Pi)
	lng2 = lng2 * (180.0 / math.Pi)

	return &Point{Lat: lat2, Lng: lng2}
}

// CalcuLates the Haversine distance between two points.
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) GreatCircleDistance(p2 *Point) float64 {
	dLat := (p2.Lat - p.Lat) * (math.Pi / 180.0)
	dLon := (p2.Lng - p.Lng) * (math.Pi / 180.0)

	lat1 := p.Lat * (math.Pi / 180.0)
	lat2 := p2.Lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}

// Calculates the initial bearing (sometimes referred to as forward azimuth)
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) BearingTo(p2 *Point) float64 {

	dLon := (p2.Lng - p.Lng) * math.Pi / 180.0

	lat1 := p.Lat * math.Pi / 180.0
	lat2 := p2.Lat * math.Pi / 180.0

	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) -
		math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)
	brng := math.Atan2(y, x) * 180.0 / math.Pi

	return brng
}
