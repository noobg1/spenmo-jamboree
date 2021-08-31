package numbersequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSequence(test *testing.T) {

	// WHEN
	resultAndExpectation := [][]bool{
		{findSequence([]int{1, 2, 5, 6}, []int{5, 6}), true},
		{findSequence([]int{1, 2, 5, 6}, []int{5, 6, 7}), false},
		{findSequence([]int{8, 7, 6, 100}, []int{5, 6, 7}), false},
		{findSequence([]int{8, 7, 6, 100}, []int{8, 7, 6}), true},
		{findSequence([]int{8, 7, 6, 100}, []int{8, 7}), true},
		{findSequence([]int{8, 7, 6, 100}, []int{}), false},
		{findSequence([]int{1, 2, 3, 4}, []int{}), false},
		{findSequence([]int{20, 7, 8, 10, 2, 5, 6}, []int{7, 8}), true},
		{findSequence([]int{20, 7, 8, 10, 2, 5, 6}, []int{8, 7}), false},
		{findSequence([]int{20, 7, 8, 10, 2, 5, 6}, []int{7, 10}), false},
	}

	// THEN
	for _, val := range resultAndExpectation {
		assert.Equal(test, val[0], val[1])
	}

}
