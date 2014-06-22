package main

import (
	"math"
	"strings"
	"reflect"
)

func contains(slice_raw interface{}, elem_raw interface{}) bool {
	elem_type := reflect.TypeOf(elem_raw)
	if reflect.TypeOf(slice_raw) != reflect.SliceOf(elem_type) {
		panic("arguments to contains are of different types!")
	}

	slice := reflect.ValueOf(slice_raw)

	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Interface() == elem_raw {
			return true
		}
	}
	return false
}

func intRoundDiv(num int, divisor int) int {
	return int((float32(num) / float32(divisor)) + .5)
}

func stringToLines(raw string) (out []string) {
	out = strings.Split(raw, "\n")
	h := len(out)
	out = out[1:h]

	return
}

func percentDiff(num int, num2 int) float64 {
	diff := math.Abs(float64(num - num2))
	avg := float64(num+num2) / 2
	return diff / avg
}
