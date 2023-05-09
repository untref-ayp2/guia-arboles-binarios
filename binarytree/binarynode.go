package binarytree

import (
	"golang.org/x/exp/constraints"
)

// BinaryNode implementa el tipo BinaryNode con dos campos de tipo
// *BinaryNode que son punteros a los hijos izquierdo y derecho, tambien de tipo BinaryNode,
// y un tercer campo de tipo T generico pero Ordered, por compatibilidad con BinarySerchTree.
type BinaryNode[T constraints.Ordered] struct {
	data  T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

// NewBinaryNode crea un nuevo BinaryNode.
//
// Uso:
//
//	d := binarytree.NewBinarNode[int](data, hIzq, hDer)
//
// Parámetros:
//   - 'data' : el dato que guarda el nodo de tipo T
//   - 'left' : el nodo que será asignado como hijo izquierdo
//   - 'right' : el nodo que será asignado como hijo derecho
//
// Retorna:
//   - un puntero a un nuevo BinaryNode.
func NewBinaryNode[T constraints.Ordered](data T, left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	// Implementar
	return nil
}

// GetData retorna el dato guardado en el nodo de tipo T.
//
// Retorna:
//   - el dato guardado en el nodo de tipo T.
func (n *BinaryNode[T]) GetData() T {
	// Implementar
	var zero T
	return zero
}

// GetLeft retorna el hijo izquierdo del nodo.
//
// Retorna:
//   - un puntero al hijo izquierdo del nodo.
func (n *BinaryNode[T]) GetLeft() *BinaryNode[T] {
	// Implementar
	return nil
}

// GetRight retorna el hijo derecho del nodo.
//
// Retorna:
//   - un puntero al hijo derecho del nodo.
func (n *BinaryNode[T]) GetRight() *BinaryNode[T] {
	// Implementar
	return nil
}

// Size evalúa el tamaño del árbol.
//
// Retorna:
//   - la cantidad de nodos del árbol
func (n *BinaryNode[T]) Size() int {
	// Implementar
	return -1
}

// Height retorna la altura del árbol.
//
// Retorna:
//   - la altura del árbol.
func (n *BinaryNode[T]) Height() int {
	// Implementar
	return -1
}
