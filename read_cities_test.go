package main

import "testing"

func TestReadCities(t *testing.T) {
	cities := readCities("cities_test.txt")
	city := cities[0]
	expected := City{"City", "123"}
	if city != expected {
		t.Error("Expected", expected, "got", city)
	}
}

