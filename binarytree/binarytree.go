// Package binarytree provee una implementación de árboles.
package binarytree

import (
	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryNode[T]
}

// NewBinaryTree crea un nuevo BinaryTree a partir de un dato de tipo T.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//
// Parámetros:
//   - 'data' : el dato de tipo T que guarda el nodo raíz
//
// Retorna:
//   - un puntero a un nuevo BinaryTree.
func NewBinaryTree[T constraints.Ordered](data T) *BinaryTree[T] {
	// Implementar
	return nil
}

// GetRoot obtiene el nodo raíz del árbol.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.GetRoot()
//
// Retorna:
//   - un puntero al nodo raíz del árbol.
func (t *BinaryTree[T]) GetRoot() *BinaryNode[T] {
	// Implementar
	return nil
}

// InsertLeft inserta del lado izquierdo de la raíz, el árbol que se pasa por parámetro
// Si el árbol está vacío lo inserta en la raíz
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertLeft(bt *BinaryTree[T]) {
	// Implementar
}

// InsertRight inserta del lado derecho de la raíz, el árbol que se pasa por parámetro
// Si el árbol está vacío lo inserta en la raíz
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertRight(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertRight(bt *BinaryTree[T]) {
	// Implementar
}

// Clear limpia el árbol poniendo la raíz en nil.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Clear()
func (t *BinaryTree[T]) Clear() {
	// Implementar
}

// IsEmpty evalúa si el árbol está vacío.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.IsEmpty()
//
// Retorna:
//   - `true` si el árbol está vacío; `false` si no lo está.
func (t *BinaryTree[T]) IsEmpty() bool {
	// Implementar
	return false
}

// Size retorna la cantidad de nodos del árbol.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Size()
//
// Retorna:
//   - la cantidad de nodos del árbol.
func (t *BinaryTree[T]) Size() int {
	// Implementar
	return -1
}

// Height etorna la altura del árbol, o sea, la distancia desde la raíz al nodo más profundo.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Height()
//
// Retorna:
//   - la altura del árbol.
func (t *BinaryTree[T]) Height() int {
	// Implementar
	return -1
}
