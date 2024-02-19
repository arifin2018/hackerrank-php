package Go

// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true

// 1, 3, 2, 6, 1, 2
// n = 6, k = 3
// 1 + 2 => 0, 2
// 1 + 2 => 0, 5
// 3 +  => 0, 5

func DivisibleSumPairs(n int32, k int32, s []int32) int32 {
	var result int32
	for key, v := range s {
		for i := key + 1; i < len(s); i++ {
			tempCount := v + s[i]
			if tempCount%k == 0 {
				result++
			}
		}
	}
	return result
}
