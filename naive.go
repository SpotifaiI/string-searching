package main

import "strings"

func NaiveSearch(text, pattern string) (int, []int) {
	text = strings.ToLower(text)
	pattern = strings.ToLower(pattern)
	count := 0
	positions := []int{}
	n := len(text)
	m := len(pattern)

	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			count++
			positions = append(positions, i)
		}
	}
	return count, positions
}
