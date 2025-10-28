package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	s "strings"
)

func wordFrequency(words string) map[string]int {
	words = s.ToLower(words)

	re := regexp.MustCompile(`\p{P}`)

	words = re.ReplaceAllString(words, "")

	wordsList := s.Fields(words)

	result := make(map[string]int)

	for _, word := range wordsList {
		result[word] += 1
	}

	return result
}

func main() {
	var sentence string
	fmt.Println("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There was an error while reading the input!")
		return
	}
	sentence = s.TrimSpace(input)
	fmt.Println(wordFrequency(sentence))
}
