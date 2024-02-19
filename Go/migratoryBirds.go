package Go

// https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true
func MigratoryBirds(arr []int32) int32 {
	var data map[int32]int32
	var result int32
	var resultValue int32
	data = map[int32]int32{}

	for i := 0; i < len(arr); i++ {
		data[arr[i]] = data[arr[i]] + 1
	}

	for k, v := range data {
		if v >= resultValue && k < result {
			resultValue = v
			result = k
		} else if v > resultValue {
			resultValue = v
			result = k
		}
	}

	return result
}
