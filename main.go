package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	swap(arr, 0, 1)

	fmt.Println(arr)
}


func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
