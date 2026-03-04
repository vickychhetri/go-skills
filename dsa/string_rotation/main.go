package main

import "fmt"

func main() {
	fmt.Println("Rotation of String")
	if isSubtring("waterbotlles", "erbotlleswat") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func isSubtring(str, str1 string) bool {
	if len(str) != len(str1) {
		return false
	}
	strN := []rune(str)
	strN1 := []rune(str1)
	j := 0
	i := 0

	count := 0
	var newString []rune
	var notmatch []int

	for i < len(strN) {
		if strN1[j] == strN[i] {
			count++
			j++
			newString = append(newString, strN[i])
		} else {
			notmatch = append(notmatch, i)
		}
		i++
	}

	for v := range notmatch {
		newString = append(newString, strN[v])
	}

	if string(newString) == str1 {
		return true
	}

	return false
}
