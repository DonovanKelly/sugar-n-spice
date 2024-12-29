package all

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6}
	assert.True(
		t,
		All[int](
			testSlice,
			func(i int) bool {
				return i < 7
			},
		))
	assert.False(
		t,
		All[int](
			testSlice,
			func(i int) bool {
				return i < 6
			},
		))
}
