package a

import (
	"fmt"
	"io/ioutil"
	"math"
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

	grid := getGrid(coordinates)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			closestIndex := -1
			var distance float64
			for index, coordinate := range coordinates {
				// find how close to cooardinate
				mD := manhattanDistance(x, y, coordinate.x, coordinate.y)
				if index == 0 {
					distance = mD
					closestIndex = index
					continue
				}

				if mD == distance {
					closestIndex = -1
					continue
				}

				if mD < distance {
					distance = mD
					closestIndex = index
				}
			}
			grid[x][y] = closestIndex
		}
	}

	dict := make(map[int]int)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == -1 {
				continue
			}
			dict[grid[x][y]] = dict[grid[x][y]] + 1
		}
	}

	var largest int
	for _, item := range dict {
		if item > largest {
			largest = item
		}
	}

	fmt.Println(largest)
}

func manhattanDistance(a int, b int, c int, d int) float64 {
	return math.Abs(float64(a)-float64(c)) + math.Abs(float64(b)-float64(d))
}

func getGridArea(coordinates []Coordinate) (int, int) {
	var maxX, maxY int
	for _, coordinate := range coordinates {
		if coordinate.x > maxX {
			maxX = coordinate.x
		}
		if coordinate.y > maxY {
			maxY = coordinate.y
		}
	}
	return maxX + 1, maxY + 1
}

func getGrid(coordinates []Coordinate) [][]int {
	maxX, maxY := getGridArea(coordinates)

	grid := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		grid[i] = make([]int, maxY)
	}
	return grid
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
