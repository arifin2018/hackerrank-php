package Go

import "math"

// https://www.hackerrank.com/challenges/cats-and-a-mouse/problem?isFullScreen=true
/*
*
Tugas Anda adalah menentukan kucing mana yang akan mencapai tikus terlebih dahulu,dengan asumsi tikus tidak bergerak dan kucing bergerak dengan kecepatan yang sama. Jika kucing datang pada saat yang bersamaan, tikus akan dibiarkan bergerak dan kabur saat mereka berkelahi.

Jika kucing A menangkap mouse terlebih dahulu, cetak `Cat A`.
Jika kucing B menangkap mouse terlebih dahulu, cetak `Cat B`.
Jika kedua kucing mencapai tikus pada saat yang sama, cetaklah `Mouse C`` saat kedua kucing tersebut berkelahi dan tikus melarikan diri.


SOAL 1:
x = 2 (Cat A)
y = 5 (Cat B)
z = 4 (Mouse C)

- jika kita liat `kucing A`, lebih jauh dari pada `kucing B` karena nilai `kucing B` mendekati `tikus C`
- carilah angka `kucing` yang mendekati `tikus C`



*/
func CatAndMouse1(x int32, y int32, z int32) string {
	var result string

	if math.Abs(float64(x-z)) < math.Abs(float64(y-z)) {
		result = "Cat A"
	} else if math.Abs(float64(y-z)) < math.Abs(float64(x-z)) {
		result = "Cat B"
	} else {
		result = "Mouse C"
	}

	return result
}
func CatAndMouse(x int32, y int32, z int32) string {
	var result string
	var resultX int32
	var resultY int32

	resultX = x - z
	resultY = y - z

	if resultX < 0 {
		resultX *= -1
	}
	if resultY < 0 {
		resultY *= -1
	}

	if resultX < resultY {
		result = "Cat A"
	} else if resultY < resultX {
		result = "Cat B"
	} else {
		result = "Mouse C"
	}

	return result
}
