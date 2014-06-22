package main

import "strings"

func contains(list []interface{}, elem interface{}) bool {
	for _, t := range list {
		if t == elem {
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
