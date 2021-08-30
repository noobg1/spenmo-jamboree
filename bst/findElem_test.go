package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElemEx0(test *testing.T) {
	// GIVEN
	treeElements := [7]int{20, 5, 6, 7, 8, 10, 40}
	tree := createBST(treeElements[:])

	// WHEN
	resultAndExpectation := [][]bool{
		{findElem(&tree, 8), true},
		{findElem(&tree, 11), false},
		{findElem(&tree, 13), false},
		{findElem(&tree, 7), true},
	}

	// THEN
	for _, val := range resultAndExpectation {
		assert.Equal(test, val[0], val[1])
	}

}

func TestFindElemEx1(test *testing.T) {
	// GIVEN
	treeElements := [7]int{11, 11, 6, 8, 101}
	tree := createBST(treeElements[:])

	// WHEN
	resultAndExpectation := [][]bool{
		{findElem(&tree, 8), true},
		{findElem(&tree, 11), true},
		{findElem(&tree, 13), false},
	}

	// THEN
	for _, val := range resultAndExpectation {
		assert.Equal(test, val[0], val[1])
	}

}
