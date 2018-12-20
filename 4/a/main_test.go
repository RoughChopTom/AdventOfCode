package a

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
	"time"
)

type Guard struct {
	asleep []time.Time
	awake []time.Time
	id string
}

func Test_Test(t *testing.T){
	content, _ := ioutil.ReadFile("testinput.txt")
	stringy := string(content)
	foo := strings.FieldsFunc(stringy, func(r rune) bool {
		return r == '\n' || r == '\r'
	})

	guardMap := make(map[string] Guard)

	sort.Strings(foo)
	var guardNumber string
	for _, item := range foo{
		first := strings.Index(item, "[")
		last := strings.Index(item, "]")

		dateTime := item[first+1: last]
		datey, _ := time.Parse("2006-01-02 15:04",dateTime)


		isGuard := strings.Contains(item, "#")
		if isGuard{
			first := strings.Index(item, "#") + 1
			last := strings.Index(item[first:], " ")
			guardNumber = item[first: first + last]

			if _, ok := guardMap[guardNumber]; !ok {
				guardMap[guardNumber] = Guard{id: guardNumber}
			}
			continue
		}
		if strings.Contains(item, "falls") {
			tempGuard := guardMap[guardNumber]
			tempGuard.asleep = append(tempGuard.asleep, datey)
			guardMap[guardNumber] = tempGuard
		}
		if strings.Contains(item, "wakes"){
			tempGuard := guardMap[guardNumber]
			tempGuard.awake = append(tempGuard.awake, datey)
			guardMap[guardNumber] = tempGuard
		}

	}
	fmt.Println(GetGuardWithMostMinutesAsleep(guardMap))
}

func GetGuardWithMostMinutesAsleep(guardMap map[string]Guard) string {

	var maxMinutes float64
	var guardId string
	for _, item := range guardMap{
		var mins float64
		for i := 0; i < len(item.asleep); i++ {
			sleep := item.asleep[i]
			wake := item.awake[i]


			//t1 := sleep.Format("2006-01-02 15:04")
			//t2 := wake.Format("2006-01-02 15:04")
			//_ = t1
			//_ = t2
			foo := wake.Sub(sleep)
			mins += foo.Minutes() - 1

		}
		if(mins > maxMinutes){
			maxMinutes = mins
			guardId = item.id
		}
	}
	return guardId
}
