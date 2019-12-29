package main

import "fmt"

type password [6]int

func NewPassword(value int) (pw password) {
	for i := len(pw) - 1; i >= 0; i-- {
		pw[i] = value % 10
		value /= 10
	}

	return
}

func checkPart1(pw password) (ok bool) {
	lastDigit := -1

	for _, digit := range pw {
		if digit < lastDigit {
			return false
		} else if digit == lastDigit {
			ok = true
		}

		lastDigit = digit
	}

	return ok
}

func checkPart2(pw password) (ok bool) {
	lastDigit := -1
	strike := 0

	for _, digit := range pw {
		if digit < lastDigit {
			return false
		} else if digit > lastDigit {
			ok = ok || strike == 1
			strike = 0
		} else {
			strike += 1
		}

		lastDigit = digit
	}

	return ok || strike == 1
}

func count(min, max int, check func(password) bool) (count int) {
	for value := min; value <= max; value++ {
		if check(NewPassword(value)) {
			count++
		}
	}

	return
}

func main() {
	min, max := 359282, 820401
	fmt.Println("Part 1:", count(min, max, checkPart1))
	fmt.Println("Part 2:", count(min, max, checkPart2))
}
