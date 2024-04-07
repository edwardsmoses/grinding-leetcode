package main

import (
	"fmt"
)

var testCases = []struct {
	word1    string
	word2    string
	expected string
}{
	{
		word1:    "abc",
		word2:    "pqr",
		expected: "apbqcr",
	},
	{
		word1:    "ab",
		word2:    "pqrs",
		expected: "apbqrs",
	},
	{
		word1:    "abcd",
		word2:    "pq",
		expected: "apbqcd",
	},
	{
		word1:    "a",
		word2:    "p",
		expected: "ap",
	},
}

func main() {
	fmt.Println("Hello, World!")

	for _, testCase := range testCases {
		result := mergeAlternately(testCase.word1, testCase.word2)
		fmt.Println("Test result: ", result)
		if result != testCase.expected {
			fmt.Println("Test failed: ", testCase.expected, result)
		} else {
			fmt.Println("Test passed: ", testCase.expected, result)
		}
	}
}

func mergeAlternately(word1 string, word2 string) string {
	fmt.Println("word1: ", word1)
	fmt.Println("word2: ", word2)

	var result string

	for i := 0; i < len(word1); i++ {
		result += string(word1[i])
		if i < len(word2) {
			result += string(word2[i])
		}
	}

	return result
}
