package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_GetCheckSum(t *testing.T){
	f, _ := os.Open("input.txt")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	checkSum := CheckSum(string(content))

	if checkSum != 12{
		t.Error("CheckSum was incorrect!")
	}
}

func CheckSum(s string) int {
	splitted := strings.FieldsFunc(s, Split)

	var twos int
	var threes int
	for _, item := range splitted{
		var charsBeen []rune
		alreadyAddedATwo := false
		alreadyAddedAThree := false
		for _, char := range item{
			if contains(charsBeen, char) {
				continue
			}
			charsBeen = append(charsBeen, char)

			numOfCharInString := strings.Count(item, string(char))

			if(!alreadyAddedATwo && numOfCharInString == 2){
				twos += 1
				alreadyAddedATwo = true
			}

			if(!alreadyAddedAThree && numOfCharInString == 3){
				threes += 1
				alreadyAddedAThree = true
			}
		}

	}


	return twos * threes
}

func Split(r rune) bool{
	return r == '\n' || r == '\r'
}

func contains(slice []rune, item rune) bool {
	set := make(map[rune]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}