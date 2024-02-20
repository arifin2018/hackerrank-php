package Go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
