package any

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5, 6}
	assert.True(
		t,
		Any[int](
			testSlice,
			func(i int) bool {
				return i == 1
			},
		))
	assert.False(
		t,
		Any[int](
			testSlice,
			func(i int) bool {
				return i == 0
			},
		))
}
