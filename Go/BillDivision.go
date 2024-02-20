package Go

import "fmt"

// https://www.hackerrank.com/challenges/bon-appetit/problem?isFullScreen=true
func removeIndex(s []int32, index int32) []int32 {
	return append(s[:index], s[index+1:]...)
}

func BonAppetit(bill []int32, k int32, b int32) {
	var countBill int32
	var data []int32
	var result int32
	removeIndex(bill, k)
	data = bill[:len(bill)-1]
	for _, v := range data {
		countBill += v
	}
	result = b - (countBill / 2)
	if result == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(result)
	}
}
