package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	intSet := make(Set[int])
	intSet.Add(1)
	intSet.Add(2)
	assert.EqualValues(t, intSet.Size(), 2)
	intSet.Add(3)
	intSet.Add(3)
	assert.EqualValues(t, intSet.Size(), 3)
	intSet.Remove(2)
	assert.EqualValues(t, intSet.Size(), 2)
	assert.False(t, intSet.Contains(2))
	assert.True(t, intSet.Contains(1))

	list := intSet.ToSlice()
	assert.Contains(t, list, 1)
	assert.Contains(t, list, 3)
}

func TestSet_Copy(t *testing.T) {
	intSet := NewSetFromSlice[int]([]int{1, 2, 3})

	copySet := intSet.Copy()

	intSet.Remove(1)
	assert.Equal(t, intSet.Size(), 2)
	assert.Equal(t, copySet.Size(), 3)

}
