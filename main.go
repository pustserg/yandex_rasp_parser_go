package main

import (
	"bufio"
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

const (
	inputCitiesFileName      string = "cities.txt"
	outputDirectionsFileName string = "directions_out.txt"
	outputStationsFileName   string = "station_out.txt"
)

func main() {
	cities := readCities(inputCitiesFileName)
	directions := getDirectionsFromCities(&cities)
	writeDirectionsOutput(directions)
	stations := getStationsFromDirections(&directions)
	writeStationsOutput(stations)
}

func readCities(fileName string) []City {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	cities := make([]City, 0)
	for _, line := range lines {
		str := strings.Split(line, ";")
		if len(str) > 1 {
			name, code := str[0], str[1]
			city := City{name: name, code: code}
			cities = append(cities, city)
		}
	}
	return cities
}

func getDirectionsFromCities(cities *[]City) []Direction {
	directions := make([]Direction, 0)
	for _, city := range *cities {
		cityDirections := parseCityPage(&city)
		directions = append(directions, cityDirections...)
	}
	return directions
}

func getStationsFromDirections(directions *[]Direction) []Station {
	stations := make([]Station, 0)
	for _, direction := range *directions {
		directionStations := parseDirectionPage(&direction)
		stations = append(stations, directionStations...)
	}
	return stations
}

func writeDirectionsOutput(directions []Direction) {
	file := openOrCreateFile(outputDirectionsFileName)
	defer file.Close()
	for _, direction := range directions {
		_, err := file.WriteString(direction.ToString() + "\n")
		if err != nil {
			panic(err)
		}
	}
	err := file.Sync()
	if err != nil {
		panic(err)
	}
}

func writeStationsOutput(stations []Station) {
	file := openOrCreateFile(outputStationsFileName)
	defer file.Close()
	for _, station := range stations {
		_, err := file.WriteString(station.ToString() + "\n")
		if err != nil {
			panic(err)
		}
	}
	err := file.Sync()
	if err != nil {
		panic(err)
	}
}

func parseCityPage(city *City) []Direction {
	doc, err := goquery.NewDocument(city.FullUrl())
	if err != nil {
		panic(err)
	}
	directions := make([]Direction, 0)
	doc.Find(".directions-menu__item").Each(
		func(i int, s *goquery.Selection) {
			directionName := s.Text()
			directionUrl, _ := s.Find("a").Attr("href")
			url := strings.Split(directionUrl, "?")[1]
			direction := Direction{name: directionName, url: url, city: *city}
			directions = append(directions, direction)
		})
	return directions
}

func parseDirectionPage(direction *Direction) []Station {
	doc, err := goquery.NewDocument(direction.FullUrl())
	if err != nil {
		panic(err)
	}
	stations := make([]Station, 0)
	doc.Find(".b-scheme__station").Each(
		func(i int, s *goquery.Selection) {
			stationName := s.Find("a").Text()
			stationUrl, _ := s.Find("a").Attr("href")
			if len(stationName) > 0 && len(stationUrl) > 0 {
				urlWithotQuery := strings.Split(stationUrl, "?")[0]
				stationCode := strings.Split(urlWithotQuery, "/")[2]
				station := Station{
					name:      stationName,
					code:      stationCode,
					direction: *direction}
				stations = append(stations, station)
			}
		})
	return stations
}

func openOrCreateFile(fileName string) *os.File {
	_, checkError := os.Stat(fileName)
	var file *os.File
	var err error
	if os.IsNotExist(checkError) {
		file, err = os.Create(fileName)
	} else {
		file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	}
	if err != nil {
		panic(err)
	}
	return file
}
