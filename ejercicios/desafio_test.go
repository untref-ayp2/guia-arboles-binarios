package ejercicios

import (
	"testing"
	"untref-ayp2/guia-arboles-binarios/binarytree"

	"github.com/stretchr/testify/assert"
)

func TestLadronDeCasasEjemplo1(t *testing.T) {
	raiz := binarytree.NewBinaryTree(3)
	a := binarytree.NewBinaryTree(2)
	b := binarytree.NewBinaryTree(3)
	c := binarytree.NewBinaryTree(3)
	d := binarytree.NewBinaryTree(1)

	raiz.InsertLeft(a)
	raiz.InsertRight(b)
	a.InsertRight(c)
	b.InsertRight(d)

	assert.Equal(t, 7, LadronDeCasas(raiz))
}

func TestLadronDeCasasEjemplo2(t *testing.T) {
	raiz := binarytree.NewBinaryTree(3)
	a := binarytree.NewBinaryTree(4)
	b := binarytree.NewBinaryTree(5)
	c := binarytree.NewBinaryTree(1)
	d := binarytree.NewBinaryTree(3)
	e := binarytree.NewBinaryTree(1)

	raiz.InsertLeft(a)
	raiz.InsertRight(b)
	a.InsertLeft(c)
	a.InsertRight(d)
	b.InsertRight(e)

	assert.Equal(t, 9, LadronDeCasas(raiz))
}
