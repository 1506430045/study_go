package sort

func BubbleSort(arr []int) []int {
	var length, i, j int
	length = len(arr)

	if length <= 1 {
		return arr
	}

	for i = 0; i < length-1; i++ {
		for j = i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
	return arr
}

func QuickSort() {

}

func SelectSort(arr []int) []int {
	var length, i, j int
	length = len(arr)

	if length <= 1 {
		return arr
	}
	for i = 0; i < length-1; i++ {
		minIndex := i
		for j = i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swap(arr, minIndex, i)
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	var length, i, j int
	length = len(arr)

	if length <= 1 {
		return arr
	}
	for i = 1; i < length; i++ {
		for j = i; j >= 0; j-- {
			if j >= 1 && arr[j-1] > arr[j] {
				swap(arr, j, j-1)
			}
		}
	}
	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	return []int{}
}
