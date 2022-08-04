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

func TestMethodOne(t *testing.T) {
	v := []values{
		{
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            2,
			expectResult: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for _, i := range v {
		res := MethodOne(i.arr, i.d)
		if !reflect.DeepEqual(res, i.expectResult) {
			t.FailNow()
		}
	}
}

func TestMethodTwo(t *testing.T) {
	v := []values{
		{
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            2,
			expectResult: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for _, i := range v {
		res := MethodTwo(i.arr, i.d)
		if !reflect.DeepEqual(res, i.expectResult) {
			t.Log(res, i.expectResult)
			t.FailNow()
		}
	}
}

func TestMethodThree(t *testing.T) {
	v := []values{
		{
			arr:          []int{1, 2, 3, 4, 5, 6},
			d:            2,
			expectResult: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for _, i := range v {
		res := MethodThree(i.arr, i.d)
		if !reflect.DeepEqual(res, i.expectResult) {
			t.Log(res, i.expectResult)
			t.FailNow()
		}
	}
}
