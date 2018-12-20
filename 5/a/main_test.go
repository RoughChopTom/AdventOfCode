package a

import (
	"fmt"
	"io/ioutil"
	"testing"
	"unicode"
)

func Test_Test(t *testing.T) {
	inputText, _ := ioutil.ReadFile("input.txt")
	isFinished := false
	for isFinished == false {
		isFinished = removeFirstPolymer(&inputText)
		foo := string(inputText)
		_ = foo
	}

	answer := len(inputText)
	fmt.Println(answer)
}

func removeFirstPolymer(modifiedText *[]byte) bool {
	foo := string(*modifiedText)
	_ = foo
	for i := 0; i < len(*modifiedText)-1; i++ {
		// if same letter (same type)
		if unicode.ToLower(rune((*modifiedText)[i])) == unicode.ToLower(rune((*modifiedText)[i+1])) {
			// is different case (opposite polarity)
			if (unicode.IsUpper(rune((*modifiedText)[i])) && unicode.IsLower(rune((*modifiedText)[i+1]))) || (unicode.IsLower(rune((*modifiedText)[i])) && unicode.IsUpper(rune((*modifiedText)[i+1]))) {
				*modifiedText = append((*modifiedText)[:i+1], (*modifiedText)[i+2:]...)
				*modifiedText = append((*modifiedText)[:i], (*modifiedText)[i+1:]...)
				return false
			}
		}
	}
	return true
}
