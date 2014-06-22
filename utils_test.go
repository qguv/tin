package main

import "testing"
import "reflect"

func TestContains(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	if !contains(s, 3) {
		t.Error("contains: slice of int should contain the int we gave it")
	}
	if contains(s, 0) {
		t.Error("contains: slice of int shouldn't contain what we didn't give it")
	}
}

func TestMaxInSlice(t *testing.T) {
	if maxInt([]int{1, 2}) != 2 {
		t.Error("maxInt: wrong answer")
	}
	if maxInt([]int{2, 1}) != 2 {
		t.Error("maxInt: wrong answer")
	}
	if maxInt([]int{-2, -1}) != -1 {
		t.Error("maxInt: wrong answer")
	}
	if maxInt([]int{-1, -2}) != -1 {
		t.Error("maxInt: wrong answer")
	}
}

func TestStringToLines(t *testing.T) {
	s := `
a test string
with multiple lines
    and four spaces on each side of this one    
`
	l := stringToLines(s)

	expected := []string{
		"a test string",
		"with multiple lines",
		"    and four spaces on each side of this one    ",
	}
	if !reflect.DeepEqual(l, expected) {
		t.Fatal("stringToLines: didn't come out the way we expected")
	}
}
