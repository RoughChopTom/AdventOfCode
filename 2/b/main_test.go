package b

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Tester(t *testing.T){
	f, _ := os.Open("input.txt")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	splitted := strings.FieldsFunc(string(content), Split)

	for _, first := range splitted{

		for _, second := range splitted{
			word := HasJustOneDifference(first, second)
			if len(word) > 0{
				return
			}
		}

	}
}

func Split(r rune) bool{
	return r == '\n' || r == '\r'
}

func HasJustOneDifference(first string, second string) string{
	differenceFound := false
	var word []byte

	for i := 0; i < len(first); i++ {
		if(first[i] != second[i]){
			if(differenceFound){
				return ""
			}
			differenceFound = true
			continue
		}
		word = append(word, first[i])
	}
	if differenceFound{
		return string(word)
	}else{
		return ""
	}
}