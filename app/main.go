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

	minLength := len(word1)
	if len(word2) < minLength {
		minLength = len(word2)
	}

	for i := 0; i < minLength; i++ {
		result += string(word1[i]) + string(word2[i])
	}

	// Append the remainder of the longer string
	if len(word1) > minLength {
		result += word1[minLength:]
	} else if len(word2) > minLength {
		result += word2[minLength:]
	}

	return result
}
