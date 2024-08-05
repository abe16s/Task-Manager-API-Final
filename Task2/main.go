package main

import (
    "fmt"
)

func main() {
	wc := WordCount("Abebe beso ? beso bela. BEso")
    fmt.Println("Word Count:", wc)

    isPalindrome := Palindrome("ab  e .bA?")
    fmt.Println("Is Palindrome:", isPalindrome)

}