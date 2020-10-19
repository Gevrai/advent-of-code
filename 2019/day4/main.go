package main

import "strconv"

func main() {
	//Puzzle input is 307237-769058

	min := 307237
	max := 769058

	count1 := 0
	count2 := 0
	for i := min; i <= max; i++ {
		s := strconv.Itoa(i)
		if validPasswordPart1(s) {
			count1++
		}
		if validPasswordPart2(s) {
			count2++
		}
	}

	println("Part 1:", count1)
	println("Part 2:", count2)
}

func validPasswordPart1(password string) bool {
	return len(password) == 6 &&
		neverDecreases(password) &&
		twoAdjacent(password)
}

func validPasswordPart2(password string) bool {
	return len(password) == 6 &&
		neverDecreases(password) &&
		twoAdjacentOnly(password)
}

func neverDecreases(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func twoAdjacent(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func twoAdjacentOnly(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			rigthIsDifferent := i+2 >= len(s) || s[i+1] != s[i+2]
			leftIsDifferent := i-1 < 0 || s[i-1] != s[i]
			if rigthIsDifferent && leftIsDifferent {
				return true
			}
		}
	}
	return false
}
