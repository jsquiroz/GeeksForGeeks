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

// MethodThree is a juggling algorithm
func MethodThree(arr []int, d int) []int {

	g_c_d := gcd(len(arr), d)

	for i := 0; i < g_c_d; i++ {
		j := i
		temp := arr[i]
		for {
			k := (j + d) % len(arr)
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}

		arr[j] = temp
	}

	return arr
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
