package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

const (
	inputCitiesFileName  string = "cities.txt"
	outputCitiesFileName string = "cities_out.txt"
	outputDirectionsFileName string = "directions_out.txt"
)

func main() {
	cities := readCities(inputCitiesFileName)
	writeCitiesOutput(cities)
	for _, city := range cities {
		directions := parseCityPage(city.FullUrl())
		writeDirectionsOutput(directions, city.Name)
	}
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
			city := City{name, code}
			cities = append(cities, city)
		}
	}
	return cities
}

func writeCitiesOutput(cities []City) {
	file := openOrCreateFile(outputCitiesFileName)
	defer file.Close()
	for _, city := range cities {
		fmt.Println(city)
		_, err := file.WriteString(city.FullUrl() + "\n")
		if err != nil {
			panic(err)
		}
	}
	err := file.Sync()
	if err != nil {
		panic(err)
	}
}

func writeDirectionsOutput(directions []Direction, cityName string) {
	file := openOrCreateFile(outputDirectionsFileName)
	defer file.Close()
	_, err := file.WriteString(cityName + "\n")
	if err != nil {
		panic(err)
	}
	for _, direction := range directions {
		_, err := file.WriteString(direction.Name + ";" + direction.Url + "\n")
		if err != nil {
			panic(err)
		}
	}
	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func parseCityPage(url string) []Direction{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	directions := make([]Direction, 0)
	doc.Find(".directions-menu__item").Each(func(i int, s *goquery.Selection) {
		directionName := s.Text()
		directionUrl, _ := s.Find("a").Attr("href")
		direction := Direction{
			directionName,
			"https://rasp.yandex.ru/" + directionUrl}
		directions = append(directions, direction)
	})
	return directions
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
