package Go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {
	test := Birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
	assert.Equal(t, int32(2), test)
}

func TestDivisibleSumPairs(t *testing.T) {
	test := DivisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	assert.Equal(t, int32(5), test)

}
