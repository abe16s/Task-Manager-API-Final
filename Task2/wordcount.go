package main

import (
	"strings"
	"regexp"
)

func WordCount(s string) map[string]int {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w\s]`)
    s = re.ReplaceAllString(s, "")

	m := make(map[string]int)
	for _, val := range strings.Fields(s) {
		m[val] += 1
	}
	return m
}