package arrgo

import (
	"testing"
)

func TestArrgoInstance(t *testing.T) {
	myArr := New[int]()

	if myArr == nil {
		t.Errorf("Error initializing the slice: %v", myArr)
	}

}

func TestArrgoPop(t *testing.T) {
	arrInt := New(1, 2, 3)

	itemPoped := arrInt.Pop()

	if itemPoped != 3 {
		t.Errorf("Expected value %d, got %d", 3, itemPoped)
	}

	expectedSlice := []int{1, 2}

	if len(*arrInt) != len(expectedSlice) {
		t.Errorf("Expected slice length %d, got %d", len(expectedSlice), len(*arrInt))
	} else {
		for i := range expectedSlice {
			if (*arrInt)[i] != expectedSlice[i] {
				t.Errorf("Expected element %d at index %d, got %d", expectedSlice[i], i, (*arrInt)[i])
			}
		}
	}
}

func TestArrgoForEach(t *testing.T) {
	arrInt := New[int]()
	arrInt.Push(1)
	arrInt.Push(2)
	arrInt.Push(3)

	var sum int

	arrInt.ForEach(func(element, index int) {
		sum += element
	})

	if sum != 6 {
		t.Errorf("Expected sum %d, got %d", 6, sum)
	}
}

func TestArrgoMap(t *testing.T) {
	arr := New(1, 2, 3, 4)

	newArr := arr.Map(func(element, index int, slice []int) int {
		return element * 2
	})

	arrLen := len(*arr)
	newArrLen := len(newArr)
	if arrLen != newArrLen {
		t.Errorf("Expected Arrgo length to be: %v, got; %v", arrLen, newArrLen)
	}
}

func TestArrgoFilter(t *testing.T) {
	arr := New(1, 2, 3, 4)
	expectArr := New(2, 4)

	filteredSlice := arr.Filter(func(element, index int, slice []int) bool {
		return element%2 == 0
	})

	if len(*expectArr) != len(filteredSlice) {
		t.Errorf("Expected slice returned length to be: %v, got; %v", len(*expectArr), len(filteredSlice))
	}

	for i, e := range filteredSlice {
		if e != (*expectArr)[i] {
			t.Errorf("Expected slice returned value in position %d to be: %v, got; %v", i, e, (*expectArr)[i])
		}
	}
}

func TestArrgoFind(t *testing.T) {
	arr := New(1, 2, 3, 4)
	expectedValue := 3

	valueFinded := arr.Find(func(element, index int, slice []int) bool {
		return element == expectedValue
	})

	if valueFinded != expectedValue {
		t.Errorf("Expected value to be: %v, got; %v", valueFinded, expectedValue)
	}
}
