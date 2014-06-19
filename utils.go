package main

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
