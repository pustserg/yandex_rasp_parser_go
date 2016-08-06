package main

import "testing"

func TestStation_ToString(t *testing.T) {
	direction := Direction{name: "dName", url: ""}
	station := Station{name: "sName", code: "1", direction: direction}
	stationString := station.ToString()
	expected := "sName;dName;1"
	if stationString != expected {
		t.Error("expected", expected, "got", stationString)
	}
}

