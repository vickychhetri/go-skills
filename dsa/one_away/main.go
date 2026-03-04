// 🧠 Algorithm (Simple)
// Compare lengths
// If difference > 1 → false
// Walk through both strings
// Allow only one mismatch
// Adjust pointer based on edit type

package main

import (
	"fmt"
	"math"
)

func OneEdit(s1, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)

	ldiff := len2 - len1
	if math.Abs(float64(ldiff)) > 1 {
		return false
	}

	i, j := 0, 0
	edited := false

	shorter, longer := s1, s2

	if len1 > len2 {
		longer, shorter = s1, s2
	}

	for i < len(shorter) && j < len(longer) {
		if shorter[i] != longer[j] {
			if edited {
				return false
			}
			edited = true
			if len1 == len2 {
				i++
			}
		} else {
			i++
		}
		j++
	}

	return true
}

func main() {
	fmt.Println("One AWAY Problem")
	r1 := OneEdit("pale", "ple") // true
	r2 := OneEdit("pale", "ale") // false
	fmt.Println(r1)
	fmt.Println(r2)
}
