package main

import (
	"math"
	"reflect"
	"strings"
)

// contains determines whether a slice contains a certain element. Elements of
// the slice must be of the same type as the testing element.
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

// maxInt returns the earliest, greatest value in a slice
func maxInt(slice []int) int {
	greatest := slice[len(slice)-1]

	for i := len(slice) - 2; i >= 0; i-- {
		this := slice[i]
		if this >= greatest {
			greatest = this
		}
	}

	return greatest
}

// stringToLines turns a string (probably a string literal) with an extra
// newline on both sides into a slice of meaningful lines. Indentation is not
// ignored.
func stringToLines(raw string) (out []string) {
	out = strings.Split(raw, "\n")
	h := len(out)
	out = out[1 : h-1]

	return
}

// intRoundDiv divides two ints as if they were floats and rounds the answer
// according to arithmetic convention.
func intRoundDiv(num int, divisor int) int {
	dividend := float32(num) / float32(divisor)
	return int(dividend + 0.5)
}

// percentDiff gives the floating-point absolute difference of two integers.
func percentDiff(num int, num2 int) float64 {
	diff := math.Abs(float64(num - num2))
	avg := float64(num+num2) / 2
	return diff / avg
}
