package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(process("один, два - это 2, три один два, много слов: один"))
	fmt.Println(process("On the 7th of April we celebrate World Health Day! " +
		"Founded in 1950 by the World Health Organization, World Health Day is " +
		"designed to focus worldwide attention to a subject of major importance " +
		"to global health each year. Health is a state of physical, mental " +
		"and social well-being, not just the absence of disease or infirmity. " +
		"A Healthy lifestyle is very important because it allows you to gain " +
		"enough energy to climb higher in other areas of your life. Be healthy " +
		"and take care of yourself"))
}

func process(str string) string {
	regexp := regexp.MustCompile("[ !,\\-:;&#%\n]")
	countMap := make(map[string]int)
	split := regexp.Split(str, -1)

	repeatedSet := make(map[string]bool)
	repeated := make([]string, 0)
	single := make([]string, 0)

	for _, s := range split {
		if s != "" {
			if _, ok := countMap[s]; ok {
				countMap[s]++
			} else {
				countMap[s] = 1
			}
		}
	}
	for _, s := range split {
		if s != "" {
			if countMap[s] > 1 {
				repeatedSet[s] = true
			} else {
				single = append(single, s)
			}
		}
	}

	for k, _ := range repeatedSet {
		repeated = append(repeated, k)
	}
	sort.Slice(repeated, func(i, j int) bool {
		return countMap[repeated[i]] > countMap[repeated[j]]
	})

	var sb strings.Builder
	for _, s := range repeated {
		sb.WriteString(fmt.Sprintf("%s(%d) ", s, countMap[s]))
	}
	for _, s := range single {
		sb.WriteString(fmt.Sprintf("%s(%d) ", s, countMap[s]))
	}
	return sb.String()
}
