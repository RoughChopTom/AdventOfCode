package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func main(){
	f, _ := os.Open("input.txt")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	hi := ""
	result := 0
	var resultList []int
	resultList = append(resultList, 0)
	doit := true
	for doit != false {
		for _, foo := range content {
			if foo == 10 {
				result += stringToInt(hi)
				if (isInSlice(resultList, result)) {
					doit = false
					return
				}
				resultList = append(resultList, result)
				hi = ""
				continue;
			}
			hi += string(foo)
		}
	}
	println(result)
}

func stringToInt(s string) int{
	i, _ := strconv.Atoi(s)
	return i
}

func isInSlice(items []int, val int) bool{
	for _, v := range items{
		if v == val{
			return true
		}
	}
	return false
}