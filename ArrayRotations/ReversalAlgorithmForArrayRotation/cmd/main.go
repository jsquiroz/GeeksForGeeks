package main

func Reversal(arr []int, key int) []int {
	if key == 0 {
		return arr
	}

	d := key % len(arr)
	reveserArray(arr, 0, d-1)
	reveserArray(arr, d, len(arr)-1)
	reveserArray(arr, 0, len(arr)-1)

	return arr
}

func reveserArray(arr []int, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		end--
		start++
	}
}
