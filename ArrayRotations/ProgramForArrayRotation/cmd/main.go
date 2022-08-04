package main

// MethodOne is using temp array
func MethodOne(arr []int, d int) []int {
	tmp := arr[0:d]
	arr = arr[d:]
	arr = append(arr, tmp...)
	return arr
}

// MethodTwo is rotate one by one
func MethodTwo(arr []int, d int) []int {

	for i := 0; i < d; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}

	return arr
}
