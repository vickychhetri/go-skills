package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "aaabcdddddeeeeeeeeeeeeeeeeeee;"

	strN := []rune(str)
	i := 0
	var newStr strings.Builder
	len1 := len(strN)
	c := 1
	for i < len1 {
		if i+1 < len1 && strN[i] == strN[i+1] {
			c++
		} else {
			if c > 1 {
				newStr.WriteRune(strN[i])
				newStr.WriteString(strconv.Itoa(c))
			} else {
				newStr.WriteRune(strN[i])
			}
			c = 1
		}
		i++
	}
	fmt.Println(newStr.String())
}
