package data

import "slices"

type Stack[T any] struct {
	Push    func(T)
	Peek    func() T
	Pop     func() T
	Length  func() int
	Reverse func()
}

func ArrToStack[T any](slice []T) Stack[T] {
	return Stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			indexOfLast := len(slice) - 1
			res := slice[indexOfLast]
			slice = slice[:indexOfLast]
			return res
		},
		Length: func() int {
			return len(slice)
		},
		Peek: func() T {
			return slice[len(slice)-1]
		},
		Reverse: func() {
			slices.Reverse(slice)
		},
	}
}

func NewStack[T any]() Stack[T] {
	return ArrToStack(make([]T, 0))
}
