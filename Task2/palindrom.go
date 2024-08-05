package main

import (
	"regexp"
	"strings"
)

func Palindrome(s string) bool{
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w]`)
    s = re.ReplaceAllString(s, "")
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l ++
		r --
	}
	return true
}