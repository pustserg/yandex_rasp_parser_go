package main

import "testing"

func TestDirectionToString(t *testing.T) {
	city := City{name: "cName", code: "1"}
	direction := Direction{name: "dName", url: "url", city: city}
	expected := "dName;cName;https://rasp.yandex.ru/city/1/direction?url"
	directionString := direction.ToString()
	if directionString != expected {
		t.Error("expected", expected, "got", directionString)
	}
}
