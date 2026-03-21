// 1. Reverse String (Unicode Safe)
// Problem: Reverse a string without breaking Unicode characters.
// Input:
// s = "hello"
// Output:
// "olleh"
// Edge Input:
// s = "héllo"
// Output:
// "olléh"
// Note: Use []rune, not []byte

//  func reverse(s string) string {
//     r := []rune(s)
//     i, j := 0, len(r)-1
//     for i < j {
//         r[i], r[j] = r[j], r[i]
//         i++
//         j--
//     }
//     return string(r)
// }

// WHY

// "hello"        → simple ASCII (1 byte per char)
// "héllo"        → 'é' is multibyte
// "你好"          → each character is multibyte
// "👋🌍"          → emojis are multibyte

package main

import "fmt"

func main() {
	var s = "héllo"

	rever := reverse(s)
	fmt.Println(rever)
}

func reverse(str string) string {
	s := []rune(str)
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
