package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "Kooooooordinata"
	test(str)

	str = "abcdefgh"
	test(str)

	str = "aaabbbbccccc"
	test(str)

	str = "aaabcd"
	test(str)

	str = "abcdddd"
	test(str)

	str = "aaabbbaaaabbbbaaaaabbbbb"
	test(str)

	str = "aaaa"
	test(str)

	str = "aaaaa"
	test(str)
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
			if count > 4 {
				compressed = append(compressed, fmt.Sprintf("#%s#", strconv.Itoa(count)))
				compressed = append(compressed, string(prev))
				count = 1
			} else {
				for j := 0; j < count; j++ {
					compressed = append(compressed, string(prev))
				}
				count = 1
			}
		}
		prev = c
	}
	if count > 4 {
		compressed = append(compressed, fmt.Sprintf("#%s#", strconv.Itoa(count)))
	} else {
		for j := 0; j < count-1; j++ {
			compressed = append(compressed, string(prev))
		}
	}
	compressed = append(compressed, string(prev))
	return strings.Join(compressed, "")
}

func decompress(str string) string {
	count := 1
	var strCount, rv string
	isCount := false
	prev := rune(str[0])
	for i := 1; i < len(str); i++ {
		c := rune(str[i])
		if prev == '#' && unicode.IsDigit(c) {
			strCount += string(c)
			isCount = true
		} else if isCount && unicode.IsDigit(c) {
			strCount += string(prev)
		} else if isCount && unicode.IsDigit(prev) && c == '#' {
			isCount = false
		}

		if !isCount {
			if prev == '#' {
				prev = c
				continue
			}
			if strCount == "" {
				for i := 0; i < count; i++ {
					rv += string(prev)
				}
				count = 1
				isCount = false
			} else {
				count, _ = strconv.Atoi(strCount)
				strCount = ""
			}
		}
		prev = c
	}
	if !isCount {
		if strCount == "" {
			for i := 0; i < count; i++ {
				rv += string(prev)
			}
		}
	}
	return rv
}

func test(str string) {
	fmt.Println("Original    ", str)
	compressed := compress(str)
	fmt.Println("Compressed  ", compressed)
	decompressed := decompress(str)
	fmt.Println("decompressed", decompressed)
	if str == decompressed {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
