package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{
		100, 2, 1, 9, 3, 8, 6, 5,
	}
	arr = InsertSort(arr)
	fmt.Println(arr)
}

func TestRelativeSortArray(t *testing.T) {
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}

	result := RelativeSortArray(arr1, arr2)
	fmt.Println(result)
}