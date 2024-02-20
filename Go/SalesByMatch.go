package Go

import (
	"math"
)

/*
*
n adalah tumpukannya
ar adalah data kaos kakinya
kembalikan lah ada berapa pasang kaos kaki yang memiliki pasangan
*/
func SockMerchant(n int32, ar []int32) int32 {
	var result int32
	var dataStackSock = map[int32]int32{}
	for _, v := range ar {
		dataStackSock[int32(v)] += 1
	}

	for _, v := range dataStackSock {
		if v > 1 {
			result += int32(math.Floor(float64(v / 2)))
		}
	}
	return result
}
