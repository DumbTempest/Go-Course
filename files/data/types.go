package data

type distanceKm float64
type distance float64

func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.60934)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(10)
	km := d.ToKm()
	km.ToMiles()

	print(d)
}
