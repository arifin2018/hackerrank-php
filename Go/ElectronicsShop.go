package Go

// https://www.hackerrank.com/challenges/electronics-shop/problem

func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var result int32 = -1

	for _, v := range keyboards {
		for _, val := range drives {
			if result < val+v {
				if val+v <= b {
					result = val + v
				}
			}
		}
	}
	return result
}
