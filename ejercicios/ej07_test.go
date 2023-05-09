package ejercicios

import (
	"testing"
	"untref-ayp2/guia-arboles-binarios/binarytree"

	"github.com/stretchr/testify/assert"
)

func TestRecorrerPorNiveles(t *testing.T) {
	a := binarytree.NewBinaryTree("A")
	b := binarytree.NewBinaryTree("B")
	c := binarytree.NewBinaryTree("C")
	d := binarytree.NewBinaryTree("D")
	e := binarytree.NewBinaryTree("E")
	f := binarytree.NewBinaryTree("F")
	g := binarytree.NewBinaryTree("G")
	h := binarytree.NewBinaryTree("H")
	i := binarytree.NewBinaryTree("I")
	j := binarytree.NewBinaryTree("J")
	k := binarytree.NewBinaryTree("K")
	l := binarytree.NewBinaryTree("L")

	a.InsertLeft(b)

	b.InsertLeft(d)
	d.InsertLeft(g)
	d.InsertRight(h)
	h.InsertLeft(k)
	k.InsertRight(l)

	a.InsertRight(c)
	c.InsertLeft(e)
	c.InsertRight(f)
	e.InsertRight(i)
	f.InsertRight(j)

	raiz := a

	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	assert.Equal(t, expected, RecorrerPorNiveles(raiz))
}
