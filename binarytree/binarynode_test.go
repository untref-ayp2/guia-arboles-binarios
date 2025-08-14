package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinaryNode(t *testing.T) {
	nodo := NewBinaryNode(5, nil, nil)

	assert.Equal(t, 5, nodo.GetData())
	assert.Nil(t, nodo.GetLeft())
	assert.Nil(t, nodo.GetRight())
}

//      4
//     / \
//    2   5
//   / \
//  1   3

func TestSize(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4

	assert.Equal(t, 5, raiz.Size())
}

func TestHeightOnEmptyTree(t *testing.T) {
	nodo := NewBinaryNode(1, nil, nil)

	assert.Equal(t, 0, nodo.Height())
}

func TestHeight(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4

	assert.Equal(t, 2, raiz.Height())
}

func TestChildren(t *testing.T) {
	izq := NewBinaryNode(1, nil, nil)
	der := NewBinaryNode(3, nil, nil)
	raiz := NewBinaryNode(2, izq, der)

	assert.Equal(t, izq, raiz.GetLeft())
	assert.Equal(t, der, raiz.GetRight())
}
