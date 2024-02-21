package Go

import "math"

// https://www.hackerrank.com/challenges/drawing-book/problem
/**
n = jumlah halaman buku
p = nomor halaman yang akan ingin di tuju
*/
func PageCountMySelf(n int32, p int32) int32 {
	var turningFrontPage int32 = 1
	var turningBackPage int32 = 1
	var result int32

	if n%2 == 0 {
		if p == 1 {
			return 0
		} else if p == n {
			return 0
		}
		for i := 2; i < int(n); i++ {
			if int32(i) == p {
				result = turningFrontPage
				break
			}
			if i%2 != 0 {
				turningFrontPage++
			}
		}
		for i := n - 1; i > 2; i-- {
			if int32(i) == p && result > turningBackPage {
				result = turningBackPage
				break
			}
			if i%2 == 0 {
				turningBackPage++
			}
		}
	} else {
		if p == 1 {
			return 0
		} else if p == n || p == n-1 {
			return 0
		}
		for i := 2; i < int(n); i++ {
			if i == int(p) {
				result = turningFrontPage
				break
			}
			if i%2 != 0 {
				turningFrontPage++
			}
		}
		for i := n - 2; i > 1; i-- {
			if i == p && result >= turningBackPage {
				result = turningBackPage
				break
			}
			if i%2 == 0 {
				turningBackPage++
			}
		}
	}

	return result
}

/*
*
-1
2 3
4 5
-67
*/
func PageCount(n int32, p int32) int32 {
	var frontCount int32
	var backCount int32
	if p == 1 || p == n {
		return 0
	}
	if n%2 != 0 {
		if p-1 == n {
			return 0
		}
	}
	frontCount = int32(math.Floor(float64(p / 2)))
	backCount = int32(math.Floor(float64(n/2 - p/2)))

	if frontCount < backCount {
		return frontCount
	} else {
		return backCount
	}
}
