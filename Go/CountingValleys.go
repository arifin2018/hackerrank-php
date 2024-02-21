package Go

// https://www.hackerrank.com/challenges/counting-valleys/problem
/*
steps = langkah pendakian
U = menanjak pendakian (/)
D = menurun langkah pendakian (\)

pendakian selalu diawali dan diakhiri di permukaan laut
[DDUUDDUDUUUD] = 1
*/
func CountingValleys(steps int32, path string) int32 {
	var result int32
	var countData int32
	var agreeincreaseCountData bool

	for i := 0; i < int(steps); i++ {
		if string(path[i]) == "D" {
			countData--
		} else if string(path[i]) == "U" {
			countData++
		}
		if countData < 0 {
			agreeincreaseCountData = true
		} else if countData > 0 {
			agreeincreaseCountData = false
		}

		if countData == 0 && agreeincreaseCountData {
			result++
		}
	}

	return result
}
