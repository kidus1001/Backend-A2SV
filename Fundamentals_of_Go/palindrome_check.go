package main

import (
	// "fmt"
	s "strings"
	"unicode"
)

func palindromeChecker(word string) bool {
	word = s.ToLower(word)
	pureWord := ""
	for i := 0; i < len(word); i++ {
		if unicode.IsLetter(rune(word[i])) {
			pureWord += string(word[i])
		}
	}
	left := 0
	right := len(pureWord) - 1

	answer := true

	for {
		if left < right {
			if pureWord[left] != pureWord[right] {
				answer = false
				break
			}
			left += 1
			right -= 1
		} else {
			break
		}
	}
	return answer
}

// func main() {
// 	var strTocheck string
// 	fmt.Println("Enter a string: ")
// 	fmt.Scan(&strTocheck)

// 	if palindromeChecker(strTocheck) {
// 		fmt.Println("The string is palindrome.")
// 	} else {
// 		fmt.Println("The string is not palindrome.")
// 	}
// }
