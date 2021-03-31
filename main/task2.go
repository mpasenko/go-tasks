package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Kooooooordinata"
	fmt.Println(compress(str))
	str = "abcdefgh"
	fmt.Println(compress(str))
	str = "aaabbbbccccc"
	fmt.Println(compress(str))
	str = "aaabcd"
	fmt.Println(compress(str))
	str = "abcdddd"
	fmt.Println(compress(str))

}

func compress(str string) string {
	count := 1
	compressed := make([]string, 0)
	prev := rune(str[0])
	for i := 1; i < len(str); i++ {
		c := rune(str[i])
		if c == prev {
			count++
		} else {
			if count > 1 {
				compressed = append(compressed, fmt.Sprintf("#%s#", strconv.Itoa(count)))
				count = 1
			}
			compressed = append(compressed, string(prev))
		}
		prev = c
	}
	if count > 1 {
		compressed = append(compressed, fmt.Sprintf("#%s#", strconv.Itoa(count)))
		count = 1
	}
	compressed = append(compressed, string(prev))
	return strings.Join(compressed, "")
}
