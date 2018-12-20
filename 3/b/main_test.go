package a

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Cloth struct {
	Id             int
	InchesFromLeft int
	InchesFromTop  int
	Width          int
	Height         int
}

func Test_First(t *testing.T) {
	f, _ := os.Open("input.txt")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	splitted := strings.FieldsFunc(string(content), Split)

	var cloths []Cloth
	for _, line := range splitted {
		newSplitted := strings.FieldsFunc(line, NewSplit)

		newCloth := Cloth{Id: StringToInt(newSplitted[0]), InchesFromLeft: StringToInt(newSplitted[1]), InchesFromTop: StringToInt(newSplitted[2]), Width: StringToInt(newSplitted[3]), Height: StringToInt(newSplitted[4])}
		cloths = append(cloths, newCloth)
	}

	array := [1000][1000]int{}

	for _, cloth := range cloths{

		for w := cloth.InchesFromLeft; w < cloth.Width + cloth.InchesFromLeft; w++ {
			for h := cloth.InchesFromTop; h < cloth.Height + cloth.InchesFromTop; h++{
				if array[w][h] != 0 {
					array[w][h] = -1
					continue
				}
				array[w][h] = cloth.Id
			}
		}
	}

	println(IsOverlapped(cloths, array))
}

func IsOverlapped(cloths []Cloth, array [1000][1000]int) int {
	for _, cloth := range cloths{
		isOverlapped := false
		for w := cloth.InchesFromLeft; w < cloth.Width + cloth.InchesFromLeft; w++ {
			for h := cloth.InchesFromTop; h < cloth.Height + cloth.InchesFromTop; h++{
				if array[w][h] == -1 {
					isOverlapped = true
				}
			}
		}
		if !isOverlapped{
			return cloth.Id
		}
	}
	return -1
}





func Split(r rune) bool {
	return r == '\n' || r == '\r'
}

func NewSplit(r rune) bool {
	return r == '#' || r == ' ' || r == '@' || r == ',' || r == ':' || r == 'x'
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		println(err)
	}
	return i
}

func SquareInches(array [1000][1000]int) int{
	result := 0
	for i := 0; i < 1000; i++{
		for ii := 0; ii < 1000; ii++{
			if array[i][ii] != -1 {
				continue
			}
			result++
		}
	}
	return result
}
