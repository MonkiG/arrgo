package main

import (
	"reflect"
	"testing"
)

func TestArrgoInstance(t *testing.T) {
	myArr := New[int]()

	if myArr == nil {
		t.Errorf("Error initializing the slice: %v", myArr)
	}

}

func TestArrgoPop(t *testing.T) {
	arrInt := New[int]()
	arrInt.Push(1)
	arrInt.Push(2)
	arrInt.Push(3)

	itemPoped := arrInt.Pop()

	if itemPoped != 3 {
		t.Errorf("Expected value %d, got %d", 3, itemPoped)
	}

	expectedSlice := []int{1, 2}
	if !reflect.DeepEqual(arrInt.s, expectedSlice) {
		t.Errorf("Expected Arrgo slice to be: %v, got; %v", expectedSlice, arrInt.s)
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
	arr := New[int]().Init(1, 2, 3, 4)

	expectedArr := New[int]().Init(2, 4, 6, 8)

	newArr := arr.Map(func(element, index int, slice []int) int {
		return element * 2
	})

	if !reflect.DeepEqual(expectedArr.s, newArr) {
		t.Errorf("Expected Arrgo slice to be: %v, got; %v", expectedArr.s, newArr)
	}
}

func TestArrgoFilter(t *testing.T) {
	arr := New[int]().Init(1, 2, 3, 4)
	expectArr := New[int]().Init(2, 4)

	filteredSlice := arr.Filter(func(element, index int, slice []int) bool {
		return element%2 == 0
	})

	if !reflect.DeepEqual(expectArr.s, filteredSlice) {
		t.Errorf("Expected Arrgo slice to be: %v, got; %v", expectArr.s, filteredSlice)
	}
}

func TestArrgoFind(t *testing.T) {
	arr := New[int]().Init(1, 2, 3, 4)
	expectedValue := 3

	valueFinded := arr.Find(func(element, index int, slice []int) bool {
		return element == expectedValue
	})

	if valueFinded != expectedValue {
		t.Errorf("Expected value to be: %v, got; %v", valueFinded, expectedValue)
	}
}
