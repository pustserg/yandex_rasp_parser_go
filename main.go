package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputCitiesFileName  string = "cities.txt"
	outputCitiesFileName string = "cities_out.txt"
)

func main() {
	cities := readCities(inputCitiesFileName)
	writeCitiesOutput(cities)
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
	_, checkError := os.Stat(outputCitiesFileName)
	var file *os.File
	var err error
	if os.IsNotExist(checkError) {
		file, err = os.Create(outputCitiesFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	} else {
		file, err = os.OpenFile(outputCitiesFileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
	for _, city := range cities {
		fmt.Println(city)
		_, err = file.WriteString(city.FullUrl() + "\n")
		if err != nil {
			panic(err)
		}
	}
	syncErr := file.Sync()
	if syncErr != nil {
		panic(err)
	}
}
