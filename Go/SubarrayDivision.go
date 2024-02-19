package Go

// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true

/**
1. panjang data array
      -    tidak kurang dari 1
      -    tidak lebih dari 100

2. - pembagian kotak nya diambil dari D
     - m adalah batas maksimal data yang harus di jumlahkan

3.
     - dan harinya (n) tidak lebih dari int 31
     -  dan bulannya (m) tidak lebih dari int 12

kesimpulan cari angka nya jumlah yang jumlahnya sama dengan (D) dengan maksimal element adalah dari (M) cari di data (S) kembalikan lah jumlah berapa banyak data yang bisa di ambil

*/

// kesimpulan cari angka nya jumlah yang jumlahnya sama dengan (D) dengan maksimal element adalah dari (M) cari di data (S) kembalikan lah jumlah berapa banyak data yang bisa di ambil

func birthday(s []int32, d int32, m int32) int32 {
	var tempTotalResult int32
	var result int32

	for k, v := range s {
		tempTotalResult = int32(v)
		for i := 1; i < int(m); i++ {
			if k+i >= len(s) {
				break
			}
			tempTotalResult += s[k+i]
		}
		if tempTotalResult == d {
			result++
			continue
		}
	}
	return result
}
