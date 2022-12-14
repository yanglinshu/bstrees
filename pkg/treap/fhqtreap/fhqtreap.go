package fhqtreap

import (
	"bstrees/pkg/errors"
	"bstrees/pkg/trait/number"
	"bstrees/pkg/treap/node"
)

type FHQTreap[T number.Number] struct {
	Root *node.TreapNode[T]
}

func New[T number.Number]() *FHQTreap[T] {
	return &FHQTreap[T]{Root: nil}
}

func Merge[T number.Number](left *node.TreapNode[T], right *node.TreapNode[T]) *node.TreapNode[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Weight < right.Weight {
		left.Right = Merge(left.Right, right)
		left.Update()
		return left
	} else {
		right.Left = Merge(left, right.Left)
		right.Update()
		return right
	}
}

func Split[T number.Number](root *node.TreapNode[T], key T) (*node.TreapNode[T], *node.TreapNode[T]) {
	if root == nil {
		return nil, nil
	}
	if root.Value <= key {
		left, right := Split(root.Right, key)
		root.Right = left
		root.Update()
		return root, right
	} else {
		left, right := Split(root.Left, key)
		root.Left = right
		root.Update()
		return left, root
	}
}

func Kth[T number.Number](root *node.TreapNode[T], k uint32) *node.TreapNode[T] {
	for root != nil {
		leftSize := uint32(0)
		if root.Left != nil {
			leftSize = root.Left.Size
		}
		if leftSize+1 == k {
			return root
		} else if leftSize+1 < k {
			k -= leftSize + 1
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

func Find[T number.Number](root *node.TreapNode[T], value T) *node.TreapNode[T] {
	for root != nil {
		if value < root.Value {
			root = root.Left
		} else if root.Value < value {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func (thisTree *FHQTreap[T]) Insert(value T) {
	left, right := Split(thisTree.Root, value)
	thisTree.Root = Merge(Merge(left, node.New(value)), right)
}

func (thisTree *FHQTreap[T]) Delete(value T) {
	left, right := Split(thisTree.Root, value)
	left, mid := Split(left, value-1)
	if mid != nil {
		mid = Merge(mid.Left, mid.Right)
	}
	thisTree.Root = Merge(Merge(left, mid), right)
}

func (thisTree *FHQTreap[T]) Contains(value T) bool {
	return Find(thisTree.Root, value) != nil
}

func (thisTree *FHQTreap[T]) Rank(value T) uint32 {
	left, right := Split(thisTree.Root, value-1)
	defer func() {
		thisTree.Root = Merge(left, right)
	}()
	if left == nil {
		return 1
	}
	return left.Size + 1
}

func (thisTree *FHQTreap[T]) Kth(k uint32) (T, error) {
	result := Kth(thisTree.Root, k)
	if result == nil {
		return T(0), errors.ErrOutOfRange
	}
	return result.Value, nil
}

func (thisTree *FHQTreap[T]) Size() uint32 {
	if thisTree.Root == nil {
		return 0
	}
	return thisTree.Root.Size
}

func (thisTree *FHQTreap[T]) Empty() bool {
	return thisTree.Root == nil
}

func (thisTree *FHQTreap[T]) Clear() {
	thisTree.Root = nil
}

func Prev[T number.Number](root *node.TreapNode[T], value T) *node.TreapNode[T] {
	var result *node.TreapNode[T] = nil
	for root != nil {
		if root.Value < value {
			result = root
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return result
}

func (thisTree *FHQTreap[T]) Prev(value T) (T, error) {
	left, right := Split(thisTree.Root, value-1)
	defer func() {
		thisTree.Root = Merge(left, right)
	}()
	result := Kth(left, left.Size)
	if result == nil {
		return T(0), errors.ErrNoPrevValue
	}
	return result.Value, nil
}

func Next[T number.Number](root *node.TreapNode[T], value T) *node.TreapNode[T] {
	var result *node.TreapNode[T] = nil
	for root != nil {
		if value < root.Value {
			result = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return result
}

func (thisTree *FHQTreap[T]) Next(value T) (T, error) {
	left, right := Split(thisTree.Root, value)
	defer func() {
		thisTree.Root = Merge(left, right)
	}()
	result := Kth(right, 1)
	if result == nil {
		return T(0), errors.ErrNoNextValue
	}
	return result.Value, nil
}
