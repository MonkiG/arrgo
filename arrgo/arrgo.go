package arrgo

type VoidCallback[T any] func(element T, index int)
type ImmutableCallback[T any] func(element T, index int, slice []T) T
type BoolCallback[T any] func(element T, index int, slice []T) bool

type Arrayable[T any] interface {
	Push(item T)                                //☑️
	Pop() T                                     //☑️
	ForEach(callback VoidCallback[T])           //☑️
	Map(callback ImmutableCallback[T]) Arrgo[T] //☑️
	Filter(callback BoolCallback[T]) Arrgo[T]   //☑️
	Find() T                                    //☑️
	// Slice()
	// Concat()
	// Reduce()
	// Includes()
	// Every()
	// Some()
}

type Arrgo[T any] []T

func (a *Arrgo[T]) Push(item T) {
	*a = append(*a, item)
}

func (a *Arrgo[T]) Pop() T {

	if len(*a) == 0 {
		var zero T
		return zero
	}

	lastItemIndex := len(*a) - 1
	item := (*a)[lastItemIndex]
	*a = (*a)[:lastItemIndex]

	return item
}

func (a *Arrgo[T]) ForEach(callback VoidCallback[T]) {
	for i, e := range *a {
		callback(e, i)
	}
}

func (a *Arrgo[T]) Map(callback ImmutableCallback[T]) []T {
	newSlice := make([]T, 0, len(*a))
	for i, e := range *a {

		newSlice = append(newSlice, callback(e, i, *a))
	}

	return newSlice
}

func (a *Arrgo[T]) Filter(callback BoolCallback[T]) []T {
	newSlice := make([]T, 0)

	for i, e := range *a {
		if callback(e, i, *a) {
			newSlice = append(newSlice, e)
		}
	}

	return newSlice
}

func (a *Arrgo[T]) Find(callback BoolCallback[T]) T {
	var zeroValue T

	for i, e := range *a {

		if callback(e, i, *a) {

			return e
		}
	}

	return zeroValue
}

func New[T any](params ...T) *Arrgo[T] {
	var arr Arrgo[T]
	if params != nil {
		arr = Arrgo[T](params)
	} else {
		arr = Arrgo[T]{}
	}

	return &arr
}
