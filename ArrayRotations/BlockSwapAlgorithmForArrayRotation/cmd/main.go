package main

func LeftRotate(arr []int, key, l int) []int {
	return LeftRotateRec(arr, 0, key, l)
}

func LeftRotateRec(arr []int, i, key, l int) []int {
	if key == 0 || l == key {
		return arr
	}

	if (l - key) == key {
		return swap(arr, i, l-key+i, key)
	}

	if key < l-key {
		arr = swap(arr, i, l-key, key)
		return LeftRotateRec(arr, i, key, l-key)
	}

	arr = swap(arr, i, key, l-key)
	return LeftRotateRec(arr, l-key+i, 2*key-l, key)
}

func swap(arr []int, fi, si, d int) []int {

	for i := 0; i < d; i++ {
		temp := arr[fi+i]
		arr[fi+i] = arr[si+i]
		arr[si+i] = temp
	}

	return arr
}
