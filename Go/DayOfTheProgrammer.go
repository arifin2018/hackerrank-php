package Go

import "fmt"

// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true

func DayOfProgrammer(year int32) string {
	var result string
	count8Month := 215

	if year < 1918 {
		if year%4 == 0 {
			count8Month += 29
		} else {
			count8Month += 28
		}
	} else {
		if year%400 == 0 {
			count8Month += 29
		} else if year%4 == 0 && year%100 != 0 {
			count8Month += 29
		} else {
			count8Month += 28
		}
	}

	countDate := 256 - count8Month
	if year == 1918 {
		countDate = 26
	}

	result = fmt.Sprintf("%v.09.%v", countDate, year)
	return result
}
