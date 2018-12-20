package a

import (
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
	"unicode"
)

func Test_Test(t *testing.T) {
	inputText, _ := ioutil.ReadFile("input.txt")


	letters := []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}


	result := make(chan int)
	//smallestLength <- len(inputText)
	var wg sync.WaitGroup
	for _, letter := range letters{
		var modifiedInputText []byte
		wg.Add(1)
		go getAnswer(inputText, letter, modifiedInputText, result)
	}

	fmt.Println(<-result)
}

func getAnswer(inputText []byte, letter rune, modifiedInputText []byte, smallestLength chan int) {
	for _, runy := range inputText {
		if rune(runy) == letter || rune(runy) == unicode.ToUpper(letter) {
			continue
		}
		modifiedInputText = append(modifiedInputText, runy)
	}
	answer := getPolymer(modifiedInputText)

	smallestLength <- answer
}

func getPolymer(inputText []byte) int{
	isFinished := false
	for isFinished == false {
		isFinished = removeFirstReaction(&inputText)
		foo := string(inputText)
		_ = foo
	}

	return len(inputText)
}

func removeFirstReaction(modifiedText *[]byte) bool {
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
