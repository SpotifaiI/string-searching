package main

import "strings"

func RabinKarp(text, pattern string) (int, []int) {
	text = strings.ToLower(text)
	pattern = strings.ToLower(pattern)
	d := 256
	q := 101
	n := len(text)
	m := len(pattern)
	count := 0
	positions := []int{}

	if m > n {
		return 0, positions
	}

	h := 1
	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}

	p := 0
	t := 0
	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(text[i])) % q
	}

	for i := 0; i <= n-m; i++ {
		if p == t {
			match := true
			for j := 0; j < m; j++ {
				if pattern[j] != text[i+j] {
					match = false
					break
				}
			}
			if match {
				count++
				positions = append(positions, i)
			}
		}
		if i < n-m {
			t = (d*(t-int(text[i])*h) + int(text[i+m])) % q
			if t < 0 {
				t += q
			}
		}
	}
	return count, positions
}
