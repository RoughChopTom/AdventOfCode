package a

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

type Coordinate struct {
	x int
	y int
}

func Test_Test(t *testing.T) {
	content, _ := ioutil.ReadFile("testinput.txt")
	coordinates := getCoordinates(strings.FieldsFunc(string(content), invalids))
	_ = coordinates
}

func getCoordinates(stringCoordinates []string) []Coordinate {
	var coordinates []Coordinate
	for _, coordinate := range stringCoordinates {
		stringCoordinate := strings.FieldsFunc(coordinate, invalidsForCoordinate)
		x, _ := strconv.Atoi(stringCoordinate[0])
		y, _ := strconv.Atoi(stringCoordinate[1])
		coordinates = append(coordinates, Coordinate{x: x, y: y})
	}
	return coordinates
}

func invalids(r rune) bool {
	return r == '\n'
}

func invalidsForCoordinate(r rune) bool {
	return !unicode.IsNumber(r)
}
