package Go

import (
	"strconv"
)

func CountBits(num uint32) int32 {
	var result int32
	data := strconv.FormatInt(int64(num), 2)
	for _, v := range data {
		num := string(v)
		if num == "1" {
			result++
		}
	}
	return result
}
