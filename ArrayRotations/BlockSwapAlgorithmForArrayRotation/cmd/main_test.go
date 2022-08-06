package main

import (
	"log"
	"reflect"
	"testing"
)

type values struct {
	arr          []int
	d            int
	expectResult []int
}

func TestRecursive(t *testing.T) {
	v := []values{
		{
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            2,
			expectResult: []int{3, 4, 5, 6, 1, 2},
		}, {
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            4,
			expectResult: []int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, i := range v {
		res := LeftRotate(i.arr, i.d, len(i.arr))
		if !reflect.DeepEqual(res, i.expectResult) {
			log.Println(res, i.expectResult)
			t.FailNow()
		}
	}
}
