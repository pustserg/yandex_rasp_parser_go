package main

type Direction struct {
	name string
	url string
}

func (direction *Direction) FullUrl(city *City) string {
	return city.FullUrl() + "/direction?" + direction.url
}