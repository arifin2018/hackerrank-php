package Go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {
	test := birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
	assert.Equal(t, int32(2), test)

}
