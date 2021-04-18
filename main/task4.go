package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	encoded := "Qp vjg 7vj qh Crtkn yg egngdtcvg Yqtnf Jgcnvj Fca! " +
		"Hqwpfgf kp da vjg Yqtnf Jgcnvj Qticpkbcvkqp, " +
		"Yqtnf Jgcnvj Fca ku fgukipgf vq hqewu yqtnfykfg cvvgpvkqp " +
		"vq c uwdlgev qh oclqt korqtvcpeg vq inqdcn jgcnvj gcej agct. " +
		"Jgcnvj ku c uvcvg qh rjaukecn, ogpvcn cpf uqekcn ygnn-dgkpi, " +
		"pqv lwuv vjg cdugpeg qh fkugcug qt kphktokva. C Jgcnvja nkhguvang " +
		"ku xgta korqtvcpv dgecwug kv cnnqyu aqw vq ickp gpqwij gpgtia vq " +
		"enkod jkijgt kp qvjgt ctgcu qh aqwt nkhg. Dg jgcnvja cpf " +
		"vcmg ectg qh aqwtugnh"
	keyWords := "7th World Health it state a"

	key, err := getKey(encoded, strings.Split(keyWords, " "))
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("key:", key)
		fmt.Println(decode(encoded, key))
	}

}

func decode(text string, key int) string {
	rv := ""
	for _, run := range text {
		decodedRun, _ := shift1(run, -key)
		rv += string(decodedRun)
	}
	return rv
}

func getKey(text string, keyWords []string) (int, error) {
	if len(keyWords) < 5 {
		return 0, errors.New("illegal arguments error")
	}
	regexp := regexp.MustCompile("[ !,\\-:;&#%\n]")
	split := regexp.Split(text, -1)
	maxLen := 1

	for _, s := range split {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	byWordLengthEncryptedMap := make(map[int][]string, maxLen)

	for _, s := range split {
		if s != "" {
			byWordLengthEncryptedMap[len(s)-1] = append(byWordLengthEncryptedMap[len(s)-1], s)
		}
	}

	found := false
	key := -1
	globalKey := -1
	keyMap := make(map[int]int)

	for _, keyWord := range keyWords {
		if keyWord == "" {
			continue
		}
		encryptedSlice := byWordLengthEncryptedMap[len(keyWord)-1]
		for _, encryptedWord := range encryptedSlice {
			for i, run := range encryptedWord {
				shifted, _ := shift1(run, -key)
				if found && shifted == rune(keyWord[i]) {
					continue
				} else {
					key = -1
					if found {
						found = false
						break
					}
				}

				skip := false // Skip in case when we bump into for example digit
				for c := 1; c < 26; c++ {
					shifted, err := shift1(run, -c)
					if err != nil {
						skip = true
						break
					}
					r := rune(keyWord[i])
					if shifted == r {
						key = c
						found = true
						break
					}
				}
				if skip {
					continue
				}
				if !found {
					break
				}
			}
			if found {
				keyMap[key]++
				found = false
			}
		}
	}

	for key, keyCount := range keyMap {
		if keyCount >= len(keyWords) {
			if globalKey == -1 {
				globalKey = key
			} else {
				return 0, errors.New("there are more than one keys were found")
			}
		}
	}
	if globalKey == -1 {
		fmt.Println(globalKey)

		return 0, errors.New("key was not found")
	}
	return globalKey, nil
}

func shift1(r rune, shift int) (rune, error) {
	if r < 'A' || r > 'z' || (r < 'a' && r > 'Z') {
		return r, errors.New("not a letter")
	}
	low := "abcdefghijklmnopqrstuvwxyz"
	high := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shift %= 26
	s := int(r)
	if s >= 97 {
		s = s + 26 - 97 + shift
		return rune(low[s%26]), nil
	} else {
		s = s + 26 - 65 + shift
		return rune(high[s%26]), nil
	}
}
