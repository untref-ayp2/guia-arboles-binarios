package ejercicios

import (
	"testing"
	"untref-ayp2/guia-arboles-binarios/binarytree"

	"github.com/stretchr/testify/assert"
)

func TestSumaPares(t *testing.T) {
	//	   4
	//	  / \
	//	 2   5
	//	/ \
	// 1   3

	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t5 := binarytree.NewBinaryTree(5)
	t4 := binarytree.NewBinaryTree(4)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)

	tree := t4

	assert.Equal(t, 6, SumaPares(tree))
}
