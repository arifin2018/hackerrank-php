package Go

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPrinter adalah mock object untuk fmt.Printer
type MockPrinter struct {
	mock.Mock
}

// Println adalah method mock untuk fmt.Printer
func (m *MockPrinter) Println(a ...interface{}) (n int, err error) {
	args := m.Called(a)
	return args.Int(0), args.Error(1)
}

//,func,TestBirthday(t,*testing.T),{
//,	test:=Birthday([]int32{1,2,1,3,2},,3,2)
//,	assert.Equal(t,int32(2),test)
//,}

//,func,TestDivisibleSumPairs(t,*testing.T),{
//,	test:=DivisibleSumPairs(6,3,[]int32{1,3,2,6,1,2})
//,	assert.Equal(t,int32(5),test)
//,}

func TestMigratoryBirds(t *testing.T) {
	test2 := MigratoryBirds([]int32{1, 4, 4, 4, 5, 3})
	assert.Equal(t, int32(4), test2)
}

func TestCountBits(t *testing.T) {
	test2 := CountBits(127)
	assert.Equal(t, int32(7), test2)
}

func TestDayOfProgrammer(t *testing.T) {
	// test := DayOfProgrammer(2017)
	// assert.Equal(t, "13.09.2017", test)
	// test2 := DayOfProgrammer(2016)
	// assert.Equal(t, "12.09.2016", test2)
	// test3 := DayOfProgrammer(2100)
	// assert.Equal(t, "13.09.2100", test3)
	// test4 := DayOfProgrammer(1800)
	// assert.Equal(t, "12.09.1800", test4)
	// test5 := DayOfProgrammer(1917)
	// assert.Equal(t, "13.09.1917", test5)
	test6 := DayOfProgrammer(1918)
	assert.Equal(t, "26.09.1918", test6)
}

func TestBonAppetit(t *testing.T) {
	// Simpan os.Stdout untuk mengembalikannya nanti
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// BonAppetit([]int32{3, 10, 2, 9}, int32(1), int32(12))
	BonAppetit([]int32{3, 10, 2, 9}, int32(1), int32(7))

	// Baca hasil dari os.Stdout
	outCh := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outCh <- buf.String()
	}()
	w.Close()
	os.Stdout = old // mengembalikan os.Stdout ke keadaan semula

	// Periksa apakah hasil sesuai dengan yang diharapkan
	result := <-outCh
	// expected := "5" + "\n"
	expected := "Bon Appetit" + "\n"
	if result != expected {
		// t.Errorf("PrintMessage(%q) => %q, want %q", "hai", result, expected)
		t.Errorf("PrintMessage(%q) => %q, want %q", "hai", result, expected)
	}

	// test2 := DayOfProgrammer(2016)
	// assert.Equal(t, "12.09.2016", test2)
	// test3 := DayOfProgrammer(2100)
	// assert.Equal(t, "13.09.2100", test3)
	// test4 := DayOfProgrammer(1800)
	// assert.Equal(t, "12.09.1800", test4)
	// test5 := DayOfProgrammer(1917)
	// assert.Equal(t, "13.09.1917", test5)
	// test6 := DayOfProgrammer(1918)
	// assert.Equal(t, "26.09.1918", test6)
}

func TestSockMerchant(t *testing.T) {
	test2 := SockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	assert.Equal(t, int32(3), test2)
}

func TestPageCount(t *testing.T) {
	// test1 := PageCount(6, 5)
	// assert.Equal(t, int32(1), test1)
	// test2 := PageCount(6, 2)
	// assert.Equal(t, int32(1), test2)
	// test3 := PageCount(5, 4)
	// assert.Equal(t, int32(0), test3)
	// test4 := PageCount(5, 3)
	// assert.Equal(t, int32(1), test4)
	// test5 := PageCount(7, 4)
	// assert.Equal(t, int32(1), test5)
	// test6 := PageCount(7, 3)
	// assert.Equal(t, int32(1), test6)
	test7 := PageCount(6, 3)
	assert.Equal(t, int32(1), test7)
}
