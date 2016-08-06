package main

type Direction struct {
	name string
	url  string
	city City
}

func (direction *Direction) FullUrl() string {
	return direction.city.FullUrl() + "/direction?" + direction.url
}

func (direction *Direction) ToString() string {
	return direction.name + ";" +
		direction.city.name + ";" +
		direction.FullUrl()
}
