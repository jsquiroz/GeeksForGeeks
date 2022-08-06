package main

import (
	"reflect"
	"testing"
)

type values struct {
	arr          []int
	d            int
	expectResult []int
}

func TestReversal(t *testing.T) {
	v := []values{
		{
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            2,
			expectResult: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for _, i := range v {
		res := Reversal(i.arr, i.d)
		if !reflect.DeepEqual(res, i.expectResult) {
			t.FailNow()
		}
	}
}
