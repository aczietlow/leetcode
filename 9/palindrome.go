package main

import (
	"fmt"
	"strconv"
)

// Given an integer x, return true if x is a palindrome, and false otherwise.
func main() {
	fmt.Println(isPalindromeWithNoString(121))
	fmt.Println(isPalindromeWithNoString(-121))
	fmt.Println(isPalindromeWithNoString(0))

}

func isPalindrome(x int) bool {
	xs := strconv.FormatInt(int64(x), 10)
	revS := []byte{}
	for i := len(xs) - 1; i >= 0; i-- {
		revS = append(revS, xs[i])
	}
	if string(revS) == xs {
		return true
	}
	return false
}

func isPalindromeWithNoString(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	h := 0
	for x > h {
		h = (h * 10) + (x % 10)
		x = x / 10
	}

	if h == x || x == h/10 {
		return true
	}
	return false
}
