package main

type VoidCallback[T any] func(element T, index int)
type ImmutableCallback[T any] func(element T, index int, slice []T) T
type BoolCallback[T any] func(element T, index int, slice []T) bool

type Arrayable[T any] interface {
	Push(item T)                           //☑️
	Pop() T                                //☑️
	ForEach(callback VoidCallback[T])      //☑️
	Map(callback ImmutableCallback[T]) []T //☑️
	Filter(callback BoolCallback[T]) []T   //☑️
	Find() T                               //☑️
	// Slice()
	// Concat()
	// Reduce()
	// Includes()
	// Every()
	// Some()
}

type Arrgo[T any] struct {
	s []T
}

func (a *Arrgo[T]) Push(item T) {
	a.s = append(a.s, item)
}

func (a *Arrgo[T]) Pop() T {

	lastItemIndex := len(a.s) - 1
	item := a.s[lastItemIndex]
	a.s = a.s[:lastItemIndex]

	return item
}

func (a *Arrgo[T]) ForEach(callback VoidCallback[T]) {
	for i, e := range a.s {
		callback(e, i)
	}
}

func (a *Arrgo[T]) Map(callback ImmutableCallback[T]) []T {
	newSlice := make([]T, 0, len(a.s))
	for i, e := range a.s {

		newSlice = append(newSlice, callback(e, i, a.s))
	}

	return newSlice
}

func (a *Arrgo[T]) Filter(callback BoolCallback[T]) []T {
	newSlice := make([]T, 0)

	for i, e := range a.s {
		if callback(e, i, a.s) {
			newSlice = append(newSlice, e)
		}
	}

	return newSlice
}

func (a *Arrgo[T]) Find(callback BoolCallback[T]) T {
	var zeroValue T

	for i, e := range a.s {

		if callback(e, i, a.s) {

			return e
		}
	}

	return zeroValue
}

func New[T any](params ...int) *Arrgo[T] {

	size, cap := 0, 0

	if len(params) > 0 {
		size = params[0]
	}

	if len(params) > 1 {
		cap = params[1]
	}

	return &Arrgo[T]{
		s: make([]T, size, cap),
	}
}

func (a *Arrgo[T]) Init(params ...T) *Arrgo[T] {
	if a == nil {
		panic("You should initialize the Arrgo object either with the function New or using the literal initiation")
	}

	a.s = append(a.s, params...)

	return a
}
