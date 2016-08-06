package main

type Station struct {
	name      string
	code      string
	direction Direction
}

func (station *Station) ToString() string {
	return station.name + ";" + station.direction.name + ";" + station.code
}
