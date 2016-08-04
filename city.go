package main

const base_url = "https://rasp.yandex.ru/city/"

type City struct {
	Name string
	code string
}

func (city *City) FullUrl() string {
	return base_url + city.code
}
