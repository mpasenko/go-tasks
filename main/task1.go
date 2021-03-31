package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	i := 25
	fmt.Println(isAutomorphic(i))
	i = 6
	fmt.Println(isAutomorphic(i))
	i = 76
	fmt.Println(isAutomorphic(i))
	i = 625
	fmt.Println(isAutomorphic(i))

}

func isAutomorphic(number int) bool {
	return bytes.HasSuffix([]byte(strconv.Itoa(number*number)), []byte((strconv.Itoa(number))))
}
