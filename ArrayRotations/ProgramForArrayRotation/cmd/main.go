package main

func MethodOne(arr []int, d int) []int {
	tmp := arr[0:d]
	arr = arr[d:]
	arr = append(arr, tmp...)
	return arr
}
