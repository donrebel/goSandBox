package algorithms

import (
	"testing"

	"github.com/testify/assert"
)

func Test_BubbleSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(input, More)
	assert.Equal(t, expectedRes, input)
}

func Test_InsertionSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	InsertionSort(input, More)
	assert.Equal(t, expectedRes, input)
}

func Test_SelectionSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SelectionSort(input, More)
	assert.Equal(t, expectedRes, input)
}

func Test_MergeSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	MergeSort(input, More)
	assert.Equal(t, expectedRes, input)
}

func Test_QuickSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSort(input, More)
	assert.Equal(t, expectedRes, input)
}

func Test_BucketSort(t *testing.T) {
	input := []int{2, 3, 6, 1, 5, 4, 9, 8, 7}
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	BucketSort(input, 1, 10)
	assert.Equal(t, expectedRes, input)
}
