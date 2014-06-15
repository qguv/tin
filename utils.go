package main

func contains(list []interface{}, elem interface{}) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}
	return false
}
